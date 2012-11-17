Web Bench in GO
===============

This started off as a discussion with a colleague to compare performance of a simple load/compute/save workload implemented in different language/frameworks.
The initial implementation he had showed go implementation at the bottom of the pile. So, I got curious and started to fiddle with it to improve perf. This is the result of that work.

Profiling
---------

1. Have a working installation of go, graphviz, ab, redis and go-redis client.
2. Checkout this project.
3. Go to profile directory and run the run.sh script.
4. After approx. 5 minutes, a couple of pdf files with profile data should be produced along with ab output in ab.log

That's it. Enjoy!

TODO
----

1. redis client seems to allocate too much memory in it's async client structure. It can be optimized by reducing the size of the request/response channels etc.
2. compress/flate code seems to do too many small memory alloc/deallocs. This causes heavy load on the GC. This can be optimized by not having to init a new zlib reader/writer for every request.

