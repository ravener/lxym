
package pages

const Ansible = "## Introduction\n\n```yaml\n---\n\"{{ Ansible }}\" is an orchestration tool written in Python.\n...\n```\n\nAnsible is (one of many) orchestration tools. It allows you to control your\nenvironment (infrastructure and code) and automate the manual tasks.\n\nAnsible has great integration with multiple operating systems (even Windows)\nand some hardware (switches, Firewalls, etc). It has multiple tools that\nintegrate with the cloud providers. Almost every noteworthy cloud provider is\npresent in the ecosystem (AWS, Azure, Google, DigitalOcean, OVH, etc...).\n\nBut ansible is way more! It provides execution plans, an API, library, and callbacks.\n\n### Main pros and cons\n\n#### Pros\n\n* It is an agent-less tool. In most scenarios, it uses ssh as a transport layer.\nIn some way you can use it as 'bash on steroids'.\n* It is very easy to start. If you are familiar with the concept of ssh - you already\nknow Ansible (ALMOST).\n* It executes 'as is' - other tools (salt, puppet, chef - might execute in\ndifferent scenario than you would expect)\n* Documentation is at the world-class standard!\n* Writing your own modules and extensions is fairly easy.\n* Ansible AWX is the open source version of Ansible Tower we have been waiting\nfor, which provides an excellent UI.\n\n#### Cons\n\n* It is an agent-less tool - every agent consumes up to 16MB ram - in some\nenvironments, it may be noticable amount.\n* It is agent-less - you have to verify your environment consistency\n'on-demand' - there is no built-in mechanism that would warn you about some\nchange automatically (this can be achieved with reasonable effort)\n* Official GUI - Ansible Tower - is great but expensive.\n* There is no 'small enterprise' payment plan, however Ansible AWX is the free\nopen source version we were all waiting for.\n\n#### Neutral\n\nMigration - Ansible <-> Salt is fairly easy - so if you would need an\nevent-driven agent environment - it would be a good choice to start quick with\nAnsible, and convert to Salt when needed.\n\n#### Some concepts\n\nAnsible uses ssh or paramiko as a transport layer. In a way you can imagine\nthat you are using a ssh with API to perform your action. The simplest way is\nto execute remote command in more controlled way (still using ssh).\nOn the other hand - in advanced scope - you can wrap Ansible (use python Ansible\ncode as a library) with your own Python scripts! It would act a\nbit like Fabric then.\n\n## Example\n\nAn example playbook to install apache and configure log level\n\n```yaml\n---\n- hosts: apache\n\n  vars:\n      apache2_log_level: \"warn\"\n\n  handlers:\n  - name: restart apache\n    service:\n      name: apache2\n      state: restarted\n      enabled: True\n    notify:\n      - Wait for instances to listen on port 80\n    become: True\n\n  - name: reload apache\n    service:\n      name: apache2\n      state: reloaded\n    notify:\n      - Wait for instances to listen on port 80\n    become: True\n\n  - name: Wait for instances to listen on port 80\n    wait_for:\n      state: started\n      host: localhost\n      port: 80\n      timeout: 15\n      delay: 5\n\n  tasks:\n  - name: Update cache\n    apt:\n      update_cache: yes\n      cache_valid_time: 7200\n    become: True\n\n  - name: Install packages\n    apt:\n      name={{ item }}\n    with_items:\n      - apache2\n      - logrotate\n    notify:\n      - restart apache\n    become: True\n\n  - name: Configure apache2 log level\n    lineinfile:\n      dest: /etc/apache2/apache2.conf\n      line: \"LogLevel {{ apache2_log_level }}\"\n      regexp: \"^LogLevel\"\n    notify:\n      - reload apache\n    become: True\n...\n```\n\n## Installation\n\n```bash\n# Universal way\n$ pip install ansible\n\n# Debian, Ubuntu\n$ apt-get install ansible\n```\n\n* [Appendix A - How do I install ansible](#infrastructure-as-a-code)\n* [Additional Reading.](http://docs.ansible.com/ansible/latest/intro_installation.html)\n\n### Your first ansible command (shell execution)\n\n```bash\n# Command pings localhost (defined in default inventory: /etc/ansible/hosts)\n$ ansible -m ping localhost\n# You should see this output\nlocalhost | SUCCESS => {\n    \"changed\": false,\n    \"ping\": \"pong\"\n}\n```\n\n### Shell Commands\n\nThere are few commands you should know about\n\n* `ansible` (to run modules in CLI)\n* `ansible-playbook` (to run playbooks)\n* `ansible-vault` (to manage secrets)\n* `ansible-galaxy` (to install roles from github/galaxy)\n\n### Module\n\nA program (usually python) that executes, does some work and returns proper\nJSON output. This program performs specialized task/action (like manage\ninstances in the cloud, execute shell command). The simplest module is called\n`ping` - it just returns a JSON with `pong` message.\n\nExample of modules:\n\n* Module: `ping` - the simplest module that is useful to verify host connectivity\n* Module: `shell` - a module that executes a shell command on a specified host(s).\n\n\n```bash\n$ ansible -m ping all\n$ ansible -m shell -a 'date; whoami' localhost #hostname_or_a_group_name\n```\n\n* Module: `command` - executes a single command that will not be processed\nthrough the shell, so variables like `$HOME` or operands like ``|` `;`` will not\nwork. The command module is more secure, because it will not be affected by the\nuser’s environment. For more complex commands - use shell module.\n\n```bash\n$ ansible -m command -a 'date; whoami' # FAILURE\n$ ansible -m command -a 'date' all\n$ ansible -m command -a 'whoami' all\n```\n\n* Module: `file` - performs file operations (stat, link, dir, ...)\n* Module: `raw` - executes a low-down and dirty SSH command, not going through\nthe module subsystem (useful to install python2.7)\n\n### Task\n\nExecution of a single Ansible **module** is called a **task**. The simplest\nmodule is called `ping` as you could see above.\n\nAnother example of the module that allows you to execute a command remotely on\nmultiple resources is called `shell`. See above how you were using them already.\n\n### Playbook\n\n**Execution plan** written in a form of script file(s) is called **playbook**.\nPlaybooks consist of multiple elements -\n* a list (or group) of hosts that 'the play' is executed against\n* `task(s)` or `role(s)` that are going to be executed\n* multiple optional settings (like default variables, and way more)\n\nPlaybook script language is YAML. You can think that playbook is very advanced\nCLI script that you are executing.\n\n#### Example of the playbook\n\nThis example-playbook would execute (on all hosts defined in inventory) two tasks:\n* `ping` that would return message *pong*\n* `shell` that execute three commands and return the output to our terminal\n\n```yaml\n- hosts: all\n\n  tasks:\n    - name: \"ping all\"\n      ping:\n\n    - name: \"execute a shell command\"\n      shell: \"date; whoami; df -h;\"\n```\n\nRun the playbook with the command:\n\n```bash\n$ ansible-playbook path/name_of_the_playbook.yml\n```\n\nNote: Example playbook is explained in the next chapter: 'Roles'\n\n### More on ansible concept\n\n### Inventory\n\nAn inventory is a set of objects or hosts, against which we are executing our\nplaybooks or single tasks via shell commands. For these few minutes, let's\nassume that we are using the default ansible inventory (which in Debian based\nsystem is placed in `/etc/ansible/hosts`).\n\n```\nlocalhost\n\n[some_group]\nhostA.mydomain.com\nhostB.localdomain\n1.2.3.4\n\n[a_group_of_a_groups:children]\nsome_group\nsome_other_group\n```\n\n* [Additional Reading.](http://docs.ansible.com/ansible/latest/intro_inventory.html)\n\n### ansible-roles (a 'template-playbooks' with right structure)\n\nYou already know that the tasks (modules) can be run via CLI. You also know the\nplaybooks - the execution plans of multiple tasks (with variables and logic).\n\nA concept called `role` was introduced for parts of the code (playbooks) that\nshould be reusable.\n\n**Role** is a structured way to manage your set of tasks, variables, handlers,\ndefault settings, and way more (meta, files, templates). Roles allow reusing\nthe same parts of code in multiple playbooks (you can parametrize the role\n'further' during its execution). Its a great way to introduce `object oriented`\nmanagement for your applications.\n\nRole can be included in your playbook (executed via your playbook).\n\n\n```yaml\n- hosts: all\n\n  tasks:\n      - name: \"ping all\"\n        ping:\n      - name: \"execute a shell command\"\n        shell: \"date; whoami; df -h;\"\n\n  roles:\n      - some_role\n      - { role: another_role, some_variable: 'learnxiny', tags: ['my_tag'] }\n\n  pre_tasks:\n      - name: some pre-task\n        shell: echo 'this task is the last, but would be executed before roles, and before tasks'\n```\n\n#### For remaining examples we would use additional repository\nThis example installs ansible in `virtualenv` so it is independent from the system.\nYou need to initialize it into your shell-context with the `source environment.sh`\ncommand.\n\nWe are going to use this repository with examples: [https://github.com/sirkubax/ansible-for-learnXinYminutes](https://github.com/sirkubax/ansible-for-learnXinYminutes)\n\n```bash\n$ # The following example contains a shell-prompt to indicate the venv and relative path\n$ git clone git@github.com:sirkubax/ansible-for-learnXinYminutes.git\nuser@host:~/$ cd ansible-for-learnXinYminutes\nuser@host:~/ansible-for-learnXinYminutes$ source environment.sh\n$\n$ # First lets execute the simple_playbook.yml\n(venv) user@host:~/ansible-for-learnXinYminutes$ ansible-playbook playbooks/simple_playbook.yml\n```\n\nRun the playbook with roles example\n\n```bash\n$ source environment.sh\n$ # Now we would run the above playbook with roles\n(venv) user@host:~/ansible-for-learnXinYminutes$ ansible-playbook playbooks/simple_role.yml\n```\n\n#### Role directory structure\n\n```\nroles/\n   some_role/\n     defaults/      # contains default variables\n     files/         # for static files\n     templates/     # for jinja templates\n     tasks/         # tasks\n     handlers/      # handlers\n     vars/          # more variables (higher priority)\n     meta/          # meta - package (role) info\n```\n\n#### Role Handlers\nHandlers are tasks that can be triggered (notified) during execution of a\nplaybook, but they execute at the very end of a playbook. It is the best way to\nrestart a service, check if the application port is active (successful\ndeployment criteria), etc.\n\nGet familiar with how you can use roles in the simple_apache_role example\n\n```\nplaybooks/roles/simple_apache_role/\n├── tasks\n│   └── main.yml\n└── templates\n    └── main.yml\n```\n\n### ansible - variables\n\nAnsible is flexible - it has 21 levels of variable precedence.\n[read more](http://docs.ansible.com/ansible/latest/playbooks_variables.html#variable-precedence-where-should-i-put-a-variable)\nFor now you should know that CLI variables have the top priority.\nYou should also know, that a nice way to pool some data is a **lookup**\n\n### Lookups\nAwesome tool to query data from various sources!!! Awesome!\nquery from:\n* pipe  (load shell command output into variable!)\n* file\n* stream\n* etcd\n* password management tools\n* url\n\n```bash\n# read playbooks/lookup.yml\n# then run\n(venv) user@host:~/ansible-for-learnXinYminutes$ ansible-playbook playbooks/lookup.yml\n```\n\nYou can use them in CLI too\n\n```yaml\nansible -m shell -a 'echo \"{{ my_variable }}\"' -e 'my_variable=\"{{ lookup(\"pipe\", \"date\") }}\"' localhost\nansible -m shell -a 'echo \"{{ my_variable }}\"' -e 'my_variable=\"{{ lookup(\"pipe\", \"hostname\") }}\"' all\n\n# Or use in playbook\n\n(venv) user@host:~/ansible-for-learnXinYminutes$ ansible-playbook playbooks/lookup.yml\n```\n\n### Register and Conditional\n\n#### Register\n\nAnother way to dynamically generate the variable content is the `register` command.\n`Register` is also useful to store an output of a task and use its value\nfor executing further tasks.\n\n```\n(venv) user@host:~/ansible-for-learnXinYminutes$ ansible-playbook playbooks/register_and_when.yml\n```\n\n```yaml\n---\n- hosts: localhost\n  tasks:\n   - name: check the system capacity\n     shell: df -h /\n     register: root_size\n\n   - name: debug root_size\n     debug:\n        msg: \"{{ root_size }}\"\n\n   - name: debug root_size return code\n     debug:\n       msg:  \"{{ root_size.rc }}\"\n\n# when: example\n\n   - name: Print this message when return code of 'check the system capacity' was ok\n     debug:\n       msg:  \"{{ root_size.rc }}\"\n     when: root_size.rc == 0\n...\n```\n\n#### Conditionals - when:\n\nYou can define complex logic with Ansible and Jinja functions. Most common is\nusage of `when:`, with some variable (often dynamically generated in previous\nplaybook steps with `register` or `lookup`)\n\n```yaml\n---\n- hosts: localhost\n  tasks:\n   - name: check the system capacity\n     shell: df -h /\n     when: some_variable in 'a string'\n  roles:\n   - { role: mid_nagios_probe, when: allow_nagios_probes }\n...\n```\n\n### ansible - tags, limit\n\nYou should know about a way to increase efficiency by this simple functionality\n\n#### TAGS\n\nYou can tag a task, role (and its tasks), include, etc, and then run only the\ntagged resources\n\n```\nansible-playbook playbooks/simple_playbook.yml --tags=tagA,tag_other\nansible-playbook playbooks/simple_playbook.yml -t tagA,tag_other\n\nThere are special tags:\n    always\n\n--skip-tags can be used to exclude a block of code\n--list-tags to list available tags\n```\n\n[Read more](http://docs.ansible.com/ansible/latest/playbooks_tags.html)\n\n#### LIMIT\n\nYou can limit an execution of your tasks to defined hosts\n\n```\nansible-playbook playbooks/simple_playbook.yml --limit localhost\n\n--limit my_hostname\n--limit groupname\n--limit some_prefix*\n--limit hostname:group #JM\n```\n\n### Templates\n\nTemplates are a powerful way to deliver some (partially) dynamic content.\nAnsible uses **Jinja2** language to describe the template.\n\n```\nSome static content\n\n{{ a_variable }}\n\n{% for item in loop_items %}\n    this line item is {{ item }}\n{% endfor %}\n```\n\nJinja may have some limitations, but it is a powerful tool that you might like.\n\nPlease examine this simple example that installs apache2 and generates\nindex.html from the template\n\"playbooks/roles/simple_apache_role/templates/index.html\"\n\n```bash\n$ source environment.sh\n$ # Now we would run the above playbook with roles\n(venv) user@host:~/ansible-for-learnXinYminutes$ ansible-playbook playbooks/simple_role.yml --tags apache2\n```\n\n#### Jinja2 CLI\n\nYou can use the jinja in the CLI too\n\n```bash\nansible -m shell -a 'echo {{ my_variable }}' -e 'my_variable=something, playbook_parameter=twentytwo' localhost\n```\n\nIn fact - jinja is used to template parts of the playbooks too\n\n```yaml\n# check part of this playbook: playbooks/roles/sys_debug/tasks/debug_time.yml\n- local_action: shell date +'%F %T'\n  register: ts\n  become: False\n  changed_when: False\n\n- name: Timestamp\n  debug: msg=\"{{ ts.stdout }}\"\n  when: ts is defined and ts.stdout is defined\n  become: False\n```\n\n#### Jinja2 filters\n\nJinja is powerful. It has many built-in useful functions.\n\n```\n# get first item of the list\n{{ some_list | first() }}\n# if variable is undefined - use default value\n{{ some_variable | default('default_value') }}\n```\n\n[Read More](http://docs.ansible.com/ansible/latest/playbooks_filters.html)\n\n### ansible-vault\n\nTo maintain **infrastructure as code** you need to store secrets. Ansible\nprovides a way to encrypt confidential files so you can store them in the\nrepository, yet the files are decrypted on-the-fly during ansible execution.\n\nThe best way to use it is to store the secret in some secure location, and\nconfigure ansible to use them during runtime.\n\n```bash\n# Try (this would fail)\n$ ansible-playbook playbooks/vault_example.yml\n\n$ echo some_very_very_long_secret > ~/.ssh/secure_located_file\n\n# in ansible.cfg set the path to your secret file\n$ vi ansible.cfg\n  ansible_vault_password_file = ~/.ssh/secure_located_file\n\n#or use env\n$ export ANSIBLE_VAULT_PASSWORD_FILE=~/.ssh/secure_located_file\n\n$ ansible-playbook playbooks/vault_example.yml\n\n  # encrypt the file\n$ ansible-vault encrypt path/somefile\n\n  # view the file\n$ ansible-vault view path/somefile\n\n  # check the file content:\n$ cat path/somefile\n\n  # decrypt the file\n$ ansible-vault decrypt path/somefile\n```\n\n### dynamic inventory\n\nYou might like to know, that you can build your inventory dynamically.\n(For Ansible) inventory is just JSON with proper structure - if you can\ndeliver that to ansible - anything is possible.\n\nYou do not need to reinvent the wheel - there are plenty of ready to use\ninventory scripts for the most popular Cloud providers and a lot of in-house\npopular usecases.\n\n[AWS example](http://docs.ansible.com/ansible/latest/intro_dynamic_inventory.html#example-aws-ec2-external-inventory-script)\n\n```bash\n$ etc/inv/ec2.py --refresh\n$ ansible -m ping all -i etc/inv/ec2.py\n```\n\n[Read more](http://docs.ansible.com/ansible/latest/intro_dynamic_inventory.html)\n\n### ansible profiling - callback\n\nPlaybook execution takes some time. It is OK. First make it run, then you may\nlike to speed things up. Since ansible 2.x there is built-in callback for task\nexecution profiling.\n\n```\nvi ansible.cfg\n# set this to:\ncallback_whitelist = profile_tasks\n```\n\n### facts-cache and ansible-cmdb\n\nYou can pull some information about your environment from another host.\nIf the information does not change - you may consider using a facts_cache\nto speed things up.\n\n```\nvi ansible.cfg\n\n# if set to a persistent type (not 'memory', for example 'redis') fact values\n# from previous runs in Ansible will be stored.  This may be useful when\n# wanting to use, for example, IP information from one group of servers\n# without having to talk to them in the same playbook run to get their\n# current IP information.\nfact_caching = jsonfile\nfact_caching_connection = ~/facts_cache\nfact_caching_timeout = 86400\n```\n\nI like to use `jsonfile` as my backend. It allows to use another project\n`ansible-cmdb` [(project on github)](https://github.com/fboender/ansible-cmdb) that generates a HTML page of your inventory\nresources. A nice 'free' addition!\n\n### Debugging ansible [chapter in progress]\n\nWhen your job fails - it is good to be effective with debugging.\n\n1. Increase verbosity by using multiple -v **[ -vvvvv]**\n2. If variable is undefined -\n`grep -R path_of_your_inventory -e missing_variable`\n3. If variable (dictionary or a list) is undefined -\n`grep -R path_of_your_inventory -e missing_variable`\n4. Jinja template debug\n5. Strange behaviour - try to run the code 'at the destination'\n\n### Infrastructure as code\n\nYou already know, that ansible-vault allows you to store your confidential data\nalong with your code. You can go further - and define your\nansible installation and configuration as code.\nSee `environment.sh` to learn how to install the ansible itself inside a\n`virtualenv` that is not attached to your operating system (can be changed by\nnon-privileged user), and as additional benefit - upgrading version of ansible\nis as easy as installing new version in new virtualenv. What is more, you can\nhave multiple versions of Ansible present at the same time.\n\n```bash\n# recreate ansible 2.x venv\n$ rm -rf venv2\n$ source environment2.sh\n\n# execute playbook\n(venv2)$ ansible-playbook playbooks/ansible1.9_playbook.yml # would fail - deprecated syntax\n\n# now lets install ansible 1.9.x next to ansible 2.x\n(venv2)$ deactivate\n$ source environment.1.9.sh\n\n# execute playbook\n(venv1.9)$ ansible-playbook playbooks/ansible1.9_playbook.yml # works!\n\n# please note that you have both venv1.9 and venv2 present - you need to (de)activate one - that is all\n```\n\n#### become-user, become\n\nIn Ansible - to become `sudo` - use the `become` parameter. Use `become_user`\nto specify the username.\n\n```\n- name: Ensure the httpd service is running\n  service:\n    name: httpd\n    state: started\n  become: true\n```\n\nNote: You may like to execute Ansible with `--ask-sudo-pass` or add the user to\nsudoers file in order to allow non-supervised execution if you require 'admin'\nprivilages.\n\n[Read more](http://docs.ansible.com/ansible/latest/become.html)\n\n## Tips and tricks\n\n#### --check -C\n\nAlways make sure that your playbook can execute in 'dry run' mode (--check),\nand its execution is not declaring 'Changed' objects.\n\n#### --diff -D\n\nDiff is useful to see nice detail of the files changed.\nIt compare 'in memory' the files like `diff -BbruN fileA fileB`.\n\n\n#### Execute hosts with 'regex'\n\n```bash\nansible -m ping web*\n```\n\n#### Host groups can be joined, negated, etc\n\n```bash\nansible -m ping web*:!backend:monitoring:&allow_change\n```\n\n#### Tagging\n\nYou should tag some (not all) objects - a task in a playbook, all tasks\nincluded form a role, etc. It allows you to execute the chosen parts of the\nplaybook.\n\n#### no_logs: True\n\nYou may see, that some roles print a lot of output in verbose mode. There is\nalso a debug module. This is the place where credentials may leak. Use `no_log`\nto hide the output.\n\n#### Debug module\n\nallows to print a value to the screen - use it!\n\n#### Register the output of a task\n\nYou can register the output (stdout), rc (return code), stderr of a task with\nthe `register` command.\n\n#### Conditionals: when:\n\n#### Loop: with, with\\_items, with\\_dict, with\\_together\n\n[Read more](http://docs.ansible.com/ansible/latest/playbooks_conditionals.html)\n\n## Additional Resources\n\n* [Servers For Hackers: An Ansible Tutorial](https://serversforhackers.com/c/an-ansible-tutorial)\n* [A system administrator's guide to getting started with Ansible - FAST!](https://www.redhat.com/en/blog/system-administrators-guide-getting-started-ansible-fast)\n* [Ansible Tower](https://www.ansible.com/products/tower) - Ansible Tower provides a web UI, dashboard and rest interface to ansible.\n* [Ansible AWX](https://github.com/ansible/awx) - The Open Source version of Ansible Tower."
