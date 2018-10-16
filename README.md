# About tornado

- FriendFeed is a relatively simple non-blocking Web server written in Python. The Web framework used by its applications looks a bit like web.py or Google's webapp, but the Web framework also contains some useful tools and optimizations to make effective use of non-blocking server environments.

- Tornado is an open-source version of our FriendFeed Web server and its common tools.

- Tornado is distinct from the mainstream Web server frameworks of today (including most Python frameworks) : it's non-blocking and quite fast. Given its non-blocking approach and the use of epoll， Tornado can handle thousands of connections per second, making Tornado an ideal framework for real-time Web services. The main purpose of developing this Web server is to handle the real-time capabilities of FriendFeed - every active user in the application of FriendFeed maintains a server connection. 

- This document was modified by vadonical in 2018, and all rights reserved.
