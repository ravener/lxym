
package pages

const Dynamic_Programming = "# Dynamic Programming\n\n## Introduction\n\nDynamic Programming is a powerful technique used for solving a particular class of problems as we will see. The idea is very simple, If you have solved a problem with the given input, then save the result for future reference, so as to avoid solving the same problem again.\n\nAlways remember!\n\"Those who can't remember the past are condemned to repeat it\"\n\n## Ways of solving such Problems\n\n1. *Top-Down* : Start solving the given problem by breaking it down. If you see that the problem has been solved already, then just return the saved answer. If it has not been solved, solve it and save the answer. This is usually easy to think of and very intuitive. This is referred to as Memoization.\n\n2. *Bottom-Up* : Analyze the problem and see the order in which the sub-problems are solved and start solving from the trivial subproblem, up towards the given problem. In this process, it is guaranteed that the subproblems are solved before solving the problem. This is referred to as Dynamic Programming.\n\n## Example of Dynamic Programming\n\nThe Longest Increasing Subsequence problem is to find the longest increasing subsequence of a given sequence. Given a sequence `S={ a1, a2, a3, a4, ............., an-1, an }` we have to find a longest subset such that for all `j` and `i`,  `j<i` in the subset `aj<ai`.\nFirst of all we have to find the value of the longest subsequences(LSi) at every index i with last element of sequence being ai. Then largest LSi would be the longest subsequence in the given sequence. To begin LSi is assigned to be one since ai is element of the sequence(Last element). Then for all `j` such that `j<i` and `aj<ai`, we find Largest LSj and add it to LSi. Then algorithm take *O(n2)* time.\n\nPseudo-code for finding the length of the longest increasing subsequence:\nThis algorithms complexity could be reduced by using better data structure rather than array. Storing predecessor array and variable like `largest_sequences_so_far` and its index would save a lot time.\n\nSimilar concept could be applied in finding longest path in Directed acyclic graph.\n\n```python\nfor i=0 to n-1\n    LS[i]=1\n    for j=0 to i-1\n        if (a[i] >  a[j] and LS[i]<LS[j])\n            LS[i] = LS[j]+1\nfor i=0 to n-1\n    if (largest < LS[i])\n```\n\n### Some Famous DP Problems\n\n- [Floyd Warshall Algorithm - Tutorial and C Program source code](http://www.thelearningpoint.net/computer-science/algorithms-all-to-all-shortest-paths-in-graphs---floyd-warshall-algorithm-with-c-program-source-code)\n- [Integer Knapsack Problem - Tutorial and C Program source code](http://www.thelearningpoint.net/computer-science/algorithms-dynamic-programming---the-integer-knapsack-problem)\n- [Longest Common Subsequence - Tutorial and C Program source code](http://www.thelearningpoint.net/computer-science/algorithms-dynamic-programming---longest-common-subsequence)\n\n## Online Resources\n\n* MIT 6.006: [Lessons 19,20,21,22](https://www.youtube.com/playlist?list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)\n* TopCoder: [Dynamic Programming from Novice to Advanced](https://www.topcoder.com/community/data-science/data-science-tutorials/dynamic-programming-from-novice-to-advanced/)\n* [CodeChef](https://www.codechef.com/wiki/tutorial-dynamic-programming)\n* [InterviewBit](https://www.interviewbit.com/courses/programming/topics/dynamic-programming/)\n* GeeksForGeeks:\n  * [Overlapping Subproblems](https://www.geeksforgeeks.org/dynamic-programming-set-1/)\n  * [Tabulation vs Memoization](https://www.geeksforgeeks.org/tabulation-vs-memoizatation/)\n  * [Optimal Substructure Property](https://www.geeksforgeeks.org/dynamic-programming-set-2-optimal-substructure-property/)\n  * [How to solve a DP problem](https://www.geeksforgeeks.org/solve-dynamic-programming-problem/)\n* [How to write DP solutions](https://www.quora.com/Are-there-any-good-resources-or-tutorials-for-dynamic-programming-DP-besides-the-TopCoder-tutorial/answer/Michal-Danilák)\n\nAnd a [quiz](https://www.commonlounge.com/discussion/cdbbfe83bcd64281964b788969247253) to test your knowledge."
