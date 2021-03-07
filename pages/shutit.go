
package pages

const Shutit = "## ShutIt\n\nShutIt is an shell automation framework designed to be easy to use.\n\nIt is a wrapper around a Python-based expect clone (pexpect).\n\nYou can look at it as 'expect without the pain'.\n\nIt is available as a pip install.\n\n## Hello World\n\nStarting with the simplest example. Create a file called example.py:\n\n```python\n\nimport shutit\nsession = shutit.create_session('bash')\nsession.send('echo Hello World', echo=True)\n```\n\nRunning this with:\n\n```bash\npython example.py\n```\n\noutputs:\n\n```bash\n$ python example.py\necho \"Hello World\"\necho \"Hello World\"\nHello World\nIans-MacBook-Air.local:ORIGIN_ENV:RhuebR2T#\n```\n\nThe first argument to 'send' is the command you want to run. The 'echo'\nargument outputs the terminal interactions. By default ShutIt is silent.\n\n'send' takes care of all the messing around with prompts and 'expects' that\nyou might be familiar with from expect.\n\n\n## Log Into a Server\n\nLet's say you want to log into a server and run a command. Change example.py\nto:\n\n```python\nimport shutit\nsession = shutit.create_session('bash')\nsession.login('ssh you@example.com', user='you', password='mypassword')\nsession.send('hostname', echo=True)\nsession.logout()\n```\n\nwhich will log you into your server (if you replace with your details) and\noutput the hostname.\n\n```\n$ python example.py\nhostname\nhostname\nexample.com\nexample.com:cgoIsdVv:heDa77HB#\n```\n\nObviously that's insecure! Instead you can run:\n\n```python\nimport shutit\nsession = shutit.create_session('bash')\npassword = session.get_input('', ispass=True)\nsession.login('ssh you@example.com', user='you', password=password)\nsession.send('hostname', echo=True)\nsession.logout()\n```\n\nwhich forces you to input the password:\n\n```\n$ python example.py\nInput Secret:\nhostname\nhostname\nexample.com\nexample.com:cgoIsdVv:heDa77HB#\n```\n\nAgain, the 'login' method handles the changing prompt from a login. You give\nShutIt the login command, the user you expect to log in as, and a password\n(if needed), and ShutIt takes care of the rest.\n\n'logout' handles the ending of a 'login', handling any changes to the prompt\nfor you.\n\n## Log Into Multiple Servers\n\nLet's say you have a server farm of two servers, and want to log onto both.\nJust create two sessions and run similar login and send commands:\n\n```python\nimport shutit\nsession1 = shutit.create_session('bash')\nsession2 = shutit.create_session('bash')\npassword1 = session1.get_input('Password for server1', ispass=True)\npassword2 = session2.get_input('Password for server2', ispass=True)\nsession1.login('ssh you@one.example.com', user='you', password=password1)\nsession2.login('ssh you@two.example.com', user='you', password=password2)\nsession1.send('hostname', echo=True)\nsession2.send('hostname', echo=True)\nsession1.logout()\nsession2.logout()\n```\n\nwould output:\n\n```bash\n$ python example.py\nPassword for server1\nInput Secret:\n\nPassword for server2\nInput Secret:\nhostname\nhostname\none.example.com\none.example.com:Fnh2pyFj:qkrsmUNs# hostname\nhostname\ntwo.example.com\ntwo.example.com:Gl2lldEo:D3FavQjA#\n```\n\n## Example: Monitor Multiple Servers\n\nWe can turn the above into a simple monitoring tool by adding some logic to\nexamine the output of a command:\n\n```python\nimport shutit\ncapacity_command=\"\"\"df / | awk '{print $5}' | tail -1 | sed s/[^0-9]//\"\"\"\nsession1 = shutit.create_session('bash')\nsession2 = shutit.create_session('bash')\npassword1 = session.get_input('Password for server1', ispass=True)\npassword2 = session.get_input('Password for server2', ispass=True)\nsession1.login('ssh you@one.example.com', user='you', password=password1)\nsession2.login('ssh you@two.example.com', user='you', password=password2)\ncapacity = session1.send_and_get_output(capacity_command)\nif int(capacity) < 10:\n\tprint('RUNNING OUT OF SPACE ON server1!')\ncapacity = session2.send_and_get_output(capacity_command)\nif int(capacity) < 10:\n\tprint('RUNNING OUT OF SPACE ON server2!')\nsession1.logout()\nsession2.logout()\n```\n\nHere you use the 'send\\_and\\_get\\_output' method to retrieve the output of the\ncapacity command (df).\n\nThere are much more elegant ways to do the above (e.g. have a dictionary of the\nservers to iterate over), but it's up to you how clever you need the Python to\nbe.\n\n\n## More Intricate IO - Expecting\n\nLet's say you have an interaction with an interactive command line application\nyou want to automate. Here we will use telnet as a trivial example:\n\n```python\nimport shutit\nsession = shutit.create_session('bash')\nsession.send('telnet', expect='elnet>', echo=True)\nsession.send('open google.com 80', expect='scape character', echo=True)\nsession.send('GET /', echo=True, check_exit=False)\nsession.logout()\n```\n\nNote the 'expect' argument. You only need to give a subset of telnet's\nprompt to match and continue.\n\nNote also the 'check\\_exit' argument in the above, which is new. We'll come back\nto that. The output of the above is:\n\n```bash\n$ python example.py\ntelnet\ntelnet> open google.com 80\nTrying 216.58.214.14...\nConnected to google.com.\nEscape character is '^]'.\nGET /\nHTTP/1.0 302 Found\nCache-Control: private\nContent-Type: text/html; charset=UTF-8\nReferrer-Policy: no-referrer\nLocation: http://www.google.co.uk/?gfe_rd=cr&ei=huczWcj3GfTW8gfq0paQDA\nContent-Length: 261\nDate: Sun, 04 Jun 2017 10:57:10 GMT\n\n<HTML><HEAD><meta http-equiv=\"content-type\" content=\"text/html;charset=utf-8\">\n<TITLE>302 Moved</TITLE></HEAD><BODY>\n<H1>302 Moved</H1>\nThe document has moved\n<A HREF=\"http://www.google.co.uk/?gfe_rd=cr&amp;ei=huczWcj3GfTW8gfq0paQDA\">\nhere\n</A>.\n</BODY></HTML>\nConnection closed by foreign host.\n```\n\nNow back to 'check\\_exit=False'. Since the telnet command returns a failure exit\ncode (1) and we don't want the script to fail, you set 'check\\_exit=False' to\nlet ShutIt know you don't care about the exit code.\n\nIf you didn't pass that argument in, ShutIt gives you an interactive terminal\nif there is a terminal to communicate with. This is called a 'pause point'.\n\n\n## Pause Points\n\nYou can trigger a 'pause point' at any point by calling\n\n```python\n[...]\nsession.pause_point('This is a pause point')\n[...]\n```\n\nwithin your script, and then continue with the script by hitting CTRL and ']'\nat the same time. This is great for debugging: add a pause point, have a look\naround, then continue. Try this:\n\n```python\nimport shutit\nsession = shutit.create_session('bash')\nsession.pause_point('Have a look around!')\nsession.send('echo \"Did you enjoy your pause point?\"', echo=True)\n```\n\nwith output like this:\n\n```bash\n$ python example.py\nHave a look around!\n\nIans-Air.home:ORIGIN_ENV:I00LA1Mq#  bash\nimiell@Ians-Air:/space/git/shutit  ⑂ master +    \nCTRL-] caught, continuing with run...\n2017-06-05 15:12:33,577 INFO: Sending:  exit\n2017-06-05 15:12:33,633 INFO: Output (squashed):  exitexitIans-Air.home:ORIGIN_ENV:I00LA1Mq#  [...]\necho \"Did you enjoy your pause point?\"\necho \"Did you enjoy your pause point?\"\nDid you enjoy your pause point?\nIans-Air.home:ORIGIN_ENV:I00LA1Mq#\n```\n\n\n## More Intricate IO - Backgrounding\n\nReturning to our 'monitoring multiple servers' example, let's imagine we\nhave a long-running task that we want to run on each server. By default, ShutIt\nworks serially which would take a long time. But we can run tasks in the\nbackground to speed things up.\n\nHere you can try an example with the trivial command: 'sleep 60'.\n\n\n```python\nimport shutit\nimport time\nlong_command=\"\"\"sleep 60\"\"\"\nsession1 = shutit.create_session('bash')\nsession2 = shutit.create_session('bash')\npassword1 = session1.get_input('Password for server1', ispass=True)\npassword2 = session2.get_input('Password for server2', ispass=True)\nsession1.login('ssh you@one.example.com', user='you', password=password1)\nsession2.login('ssh you@two.example.com', user='you', password=password2)\nstart = time.time()\nsession1.send(long_command, background=True)\nsession2.send(long_command, background=True)\nprint('That took: ' + str(time.time() - start) + ' seconds to fire')\nsession1.wait()\nsession2.wait()\nprint('That took: ' + str(time.time() - start) + ' seconds to complete')\n```\n\nMy laptop says it took 0.5 seconds to run fire those two commands, and then just\nover a minute to complete (using the 'wait' method).\n\nAgain, this is trivial, but imagine you have hundreds of servers to manage like\nthis and you can see the power it can bring in a few lines of code and one\nPython import.\n\n\n## Learn More\n\nThere's a lot more that can be done with ShutIt.\n\nTo learn more, see:\n\n[ShutIt](https://ianmiell.github.io/shutit/)\n[GitHub](https://github.com/ianmiell/shutit/blob/master/README.md)\n\nIt's a broader automation framework, and the above is its 'standalone mode'.\n\nFeedback, feature requests, 'how do I?'s highly appreciated! Reach me at\n[@ianmiell](https://twitter.com/ianmiell)"
