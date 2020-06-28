### concurrence problem

The code makes four concurrent web requests. But, as we've seen before, the race condition causes the program to exit before we have a response from them.

1. You have to wait until all request are done
2. Delete a character from one of the URLs. implement a way to cancel all the requests that are not completed yet when the error appears
