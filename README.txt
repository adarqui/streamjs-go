
	STREAMJS-GO.

	streamjs-go is a port of stream.js (http://streamjs.org). This library provides lazy evaluation via "streams" (lazy lists) for GoLang. Streamjs's original README is provided further below with the golang specific code/excerpts mixed in.

	Usage:

		git clone https://github.com/adarqui/streamjs-go
		(cd streamjs-go; make; make test)

		go get github.com/adarqui/streamjs-go
		import (
			"github.com/adarqui/streamjs-go"
		)

		s := stream.Range(1, -1)

		...



   TODO:

	Go's "lack of generics" can cause some nutty type/switch casting. I plan on making all of the streamjs-go functions more 'generic' by creating xyzBy functions which take various parameters that you can supply to operate on the stream data. This will allow me to remove all of the interface{} typing from the actual stream functions, which will clean everything up significantly. Instead, i'd write the interface{} specific functions as example "generic functions" that will be used if the requested ByFN is nil/not supplied. This will pro-up the lib.

	I don't plan on adding anything new to streamjs-go. I'd like it to remain a direct port of streamjs. Instead, i'd like to incorporate/import streamjs-go into Data.List-go. That's where I'll include the kitchen sink.

	Apologies if this README gets absolutely wrecked (formatting-wise). I saved it via links and i'm editing it in vi.

	-- adarqui



	ORIGINAL STREAMJS README WITH GOLANG SPECIFIC CODE SUBSTITUTED FOR JAVASCRIPT CODE.


   Fork me on GitHub

                                     stream.js

   stream.js is a tiny stand-alone Javascript library that unlocks a new data
   structure for you: streams.

                 <script src='stream-min.js'></script>
            

   Download stream.js v1.1, 2Kb minified

What are streams?

   Streams are an easy to use data structure similar to arrays and linked lists,
   but with some extraordinary capabilities.

What's so special about them?

   Unlike arrays, streams are a magical data structure. They can have an infinite
   number of elements. Yes, you heard right. Their power comes from being lazily
   evaluated, which in simple terms means that they can contain infinite items.

Getting started

   If you devote 10 minutes of your time to read this page, it may completely
   change the way you think about programming (unless you come from a functional
   programming background that is!). Please bear with me and let me first
   introduce you to the basic operations streams support which are similar to
   arrays and linked lists. And then we'll talk about their interesting
   properties.

   Streams are containers. They contain items. You can make a stream with some
   items using Stream.make. Just pass it as arguments the items you want to be
   part of your stream:

				js:

                 var s = Stream.make( 10, 20, 30 ); // s is now a stream containing 10, 20, and 30

				go:

                 s := Make( 10, 20, 30 ); // s is now a stream containing 10, 20, and 30
            

   Easy enough, s is now a stream containing three items: 10, 20, and 30; in this
   order. We can look at the length of the stream using s.length() and retrieve
   particular items by index using s.item( i ). The first item of the stream can
   also be obtained by calling s.head(). Let's see it in action:

				js:

                 var s = Stream.make( 10, 20, 30 );
                 console.log( s.length() );  // outputs 3
                 console.log( s.head() );    // outputs 10
                 console.log( s.item( 0 ) ); // exactly equivalent to the line above
                 console.log( s.item( 1 ) ); // outputs 20
                 console.log( s.item( 2 ) ); // outputs 30


				go:

				s = Make(10, 20, 30)
				fmt.Printf("\ts.Length() = %d\n", s.Length1())
				fmt.Printf("\ts.Head() = %d\n", s.Head1())
				fmt.Printf("\ts.Item1(0) = %d\n", s.Item1(0))
				fmt.Printf("\ts.Item1(1) = %d\n", s.Item1(1))
				fmt.Printf("\ts.Item1(2) = %d\n", s.Item1(2))

            

   The stream.js library is already loaded on this page. Feel free to fire up
   your browser's Javascript console if you want to run some of the examples or
   write your own.

   Continuing, the empty stream can be constructed either using new Stream() or
   just Stream.make(). The stream containing all the items of the original stream
   except the head can be obtained using s.tail(). Calling s.head() or s.tail()
   on an empty stream yields an error. You can check if a stream is empty using
   s.empty() which returns either true or false.

				js:

                 var s = Stream.make( 10, 20, 30 );
                 var t = s.tail();         // returns the stream that contains two items: 20 and 30
                 console.log( t.head() );  // outputs 20
                 var u = t.tail();         // returns the stream that contains one item: 30
                 console.log( u.head() );  // outputs 30
                 var v = u.tail();         // returns the empty stream
                 console.log( v.empty() ); // prints true


				go:

				s = Make(10, 20, 30)
				t := s.Tail1()
				fmt.Printf("\ts.Tail() = t.Head() = %d\n", t.Head1())
				u := t.Tail1()
				fmt.Printf("\tt.Tail() = u.Head() = %d\n", u.Head1())
				v := u.Tail1()
				fmt.Printf("\tv.Empty() = %v\n", v.Empty())


   Here's a way to print all the elements in a stream:

				js:

                 var s = Stream.make( 10, 20, 30 );
                 while ( !s.empty() ) {
                     console.log( s.head() );
                     s = s.tail();
                 }


				go:

				s = Make(10, 20, 30)
				for ; !s.Empty() ; s = s.Tail1() {
					fmt.Printf("\tValue = %d\n", s.Head1())
				}

            

   There's a convenient shortcut for that: s.print() shows all the items in your
   stream.

What else can I do with them?

   One of the useful shortcuts is the Stream.range( min, max ) function. It
   returns a stream with the natural numbers ranging from min to max inclusive.

				js:

                 var s = Stream.range( 10, 20 );
                 s.print(); // prints the numbers from 10 to 20


				go:

				s = Range(10, 20)
				s.Print(-1)

            

   You can use map, filter, and walk on your streams. s.map( f ) takes an
   argument f, a function, and runs f on every element of the stream; it returns
   the stream of the return values of that function. So you can use it to, for
   example, double the numbers in your stream:

				js:

                 function doubleNumber( x ) {
                     return 2 * x;
                 }

                 var numbers = Stream.range( 10, 15 );
                 numbers.print(); // prints 10, 11, 12, 13, 14, 15
                 var doubles = numbers.map( doubleNumber );
                 doubles.print(); // prints 20, 22, 24, 26, 28, 30


				go:

				doubleNumber := func(x int) int {
					return 2 * x
				}
				doubleNumberInterface := func(x interface{}) interface{} {
					switch v := x.(type) {
						case int:
							return doubleNumber(v)
						default:
							return v
					}
				}

				fmt.Println("\tNumbers 10-15")
				numbers := Range(10, 15)
				numbers.Print(-1)

				fmt.Println("\tNumbers 10-15 doubled")
				doubles := numbers.Map(doubleNumberInterface)
				doubles.Print(-1)
            

   Cool, right? Similarly s.filter( f ) takes an argument f, a function, and runs
   f on every element of the stream; it then returns the stream containing only
   the elements for which f returned true. So you can use it to only keep certain
   elements in your stream. Let's construct a stream keeping only the odd numbers
   of an original stream using this idea:

				js:

                 function checkIfOdd( x ) {
                     if ( x % 2 == 0 ) {
                         // even number
                         return false;
                     }
                     else {
                         // odd number
                         return true;
                     }
                 }
                 var numbers = Stream.range( 10, 15 );
                 numbers.print();  // prints 10, 11, 12, 13, 14, 15
                 var onlyOdds = numbers.filter( checkIfOdd );
                 onlyOdds.print(); // prints 11, 13, 15


				go:

				checkIfOdd := func(x interface{}) bool {
					switch v := x.(type) {
						case int:
							if (v % 2) == 0 {
								return false
							} else {
								return true
							}
						default:
							return false
					}
				}

				fmt.Println("\tNumbers 10-15")
				numbers = Range(10, 15)
				numbers.Print(-1)

				fmt.Println("\tNumbers 10-15 Filtered for odd's")
				onlyOdds := numbers.Filter(checkIfOdd)
				onlyOdds.Print(-1)


            

   Useful, yes? Finally s.walk( f ) takes an argument f, a function, and runs f
   on every element of the stream, but it doesn't affect the stream in any way.
   Here's another way to print the elements of stream:

				js:

                 function printItem( x ) {
                     console.log( 'The element is: ' + x );
                 }
                 var numbers = Stream.range( 10, 12 );
                 // prints:
                 // The element is: 10
                 // The element is: 11
                 // The element is: 12
                 numbers.walk( printItem );


				go:

				printItem := func(x interface{}) interface{} {
					fmt.Printf("The element is: %v\n", x)
					return x
				}
				numbers = Range(10, 12)
				numbers.Walk(printItem)

            

   One more useful function: s.take( n ) returns a stream with the first n
   elements of your original stream. That's useful for slicing streams:

				js:

                 var numbers = Stream.range( 10, 100 ); // numbers 10...100
                 var fewerNumbers = numbers.take( 10 ); // numbers 10...19
                 fewerNumbers.print();


				go:

				numbers = Range(10, 100)
				fewerNumbers := numbers.Take(10)
				fewerNumbers.Print(10)

            

   A few other useful things: s.scale( factor ) multiplies every element of your
   stream by factor; and s.add( t ) adds each element of the stream s to each
   element of the stream t and returns the result. Let's see an example of this:

				js:

                 var numbers = Stream.range( 1, 3 );
                 var multiplesOfTen = numbers.scale( 10 );
                 multiplesOfTen.print(); // prints 10, 20, 30
                 numbers.add( multiplesOfTen ).print(); // prints 11, 22, 33


				go:

				numbers = Range(1, 3)
				fmt.Println("\tNumbers 1-3 scaled by 10")
				multiplesOfTen := numbers.Scale(10)
				multiplesOfTen.Print(-1)
				fmt.Println("\tNumbers 1-3 scaled by 10, Add")
				numbers.Add(multiplesOfTen).Print(-1)

            

   Although we've only seen streams of numbers until now, you can also have
   streams of anything: strings, booleans, functions, objects; even arrays or
   other streams. Please note however that your streams may not contain the
   special value undefined as an item.

Show me the magic!

   Let's now start playing with infinity. Streams don't need to have a finite
   number of elements. For example, you can omit the second argument to
   Stream.range( low, high ) and write Stream.range( low ); in that case, there
   is no upper bound, and so the stream contains all the natural numbers from low
   and up. You can also omit low and it defaults to 1. In that case
   Stream.range() returns the stream of natural numbers.

Does that require infinite memory/time/processing power?

   No, it doesn't. That's the awesome part. You can run these things and they
   work really fast, like regular arrays. Here's an example that prints the
   numbers from 1 to 10:

				js:

                 var naturalNumbers = Stream.range(); // returns the stream containing all natural numbers from 1 and up
                 var oneToTen = naturalNumbers.take( 10 ); // returns the stream containing the numbers 1...10
                 oneToTen.print();


				go:

				naturalNumbers := Range(1, -1)
				oneToTen := naturalNumbers.Take(10)
				oneToTen.Print(-1)

            

You're cheating

   Yes, I am. The point is that you can think of these structures as infinite,
   and this introduces a new programming paradigm that yields concise code that
   is easy to understand and closer to mathematics than usual imperative
   programming. The library itself is very short; it's thinking about these
   concepts that matters. Let's play with this a little more and construct the
   streams containing all even numbers and all odd numbers respectively.

				js:

                 var naturalNumbers = Stream.range(); // naturalNumbers is now 1, 2, 3, ...
                 var evenNumbers = naturalNumbers.map( function ( x ) {
                     return 2 * x;
                 } ); // evenNumbers is now 2, 4, 6, ...
                 var oddNumbers = naturalNumbers.filter( function ( x ) {
                     return x % 2 != 0;
                 } ); // oddNumbers is now 1, 3, 5, ...
                 evenNumbers.take( 3 ).print(); // prints 2, 4, 6
                 oddNumbers.take( 3 ).print(); // prints 1, 3, 5


				go:

				naturalNumbers = Range(1, -1)
				evenNumbers := naturalNumbers.Map(func(x interface{}) interface{} {
					switch v := x.(type) {
						case int:
							return 2 * v
						default:
							return x
					}
				})
				oddNumbers := naturalNumbers.Filter(func(x interface{}) bool {
					switch v := x.(type) {
						case int:
							return (v % 2) != 0
						default:
							return false
					}
				})
				fmt.Println("\tTake(3) of Even Numbers from 1 to Infiniti")
				evenNumbers.Take(3).Print(-1)
				fmt.Println("\tTake(3) of Odd Numbers from 1 to Infinity")
				oddNumbers.Take(3).Print(-1)

            

   Cool, right? I kept my promise that streams are more powerful than arrays.
   Now, bear with me for a few more minutes and let's introduce a few more things
   about streams. You can make your own stream objects using new Stream() to
   create an empty stream, or new Stream( head, functionReturningTail ) to create
   a non-empty stream. In case of a non-empty stream, the first parameter is the
   head of your desired stream, while the second parameter is a function
   returning the tail (a stream with all the rest of the elements), which could
   potentially be the empty stream. Confusing? Let's look at an example:

				js:

                 var s = new Stream( 10, function () {
                     return new Stream();
                 } );
                 // the head of the s stream is 10; the tail of the s stream is the empty stream
                 s.print(); // prints 10
                 var t = new Stream( 10, function () {
                     return new Stream( 20, function () {
                         return new Stream( 30, function () {
                             return new Stream();
                         } );
                     } );
                 } );
                 // the head of the t stream is 10; its tail has a head which is 20 and a tail which
                 // has a head which is 30 and a tail which is the empty stream.
                 t.print(); // prints 10, 20, 30


				go:

				s = NewStream(10, func(v interface{}, fn STREAMFN) *Stream {
					return NewStream(nil, nil)
				})
				fmt.Printf("\ts.Head() = %d\n", s.Head1())
				s.Print(-1)

				t = NewStream(10, func(v interface{}, fn STREAMFN) *Stream {
					return NewStream(20, func(v interface{}, fn STREAMFN) *Stream {
						return NewStream(30, func(v interface{}, fn STREAMFN) *Stream {
							return NewStream(nil, nil)
						})
					})
				})

				fmt.Println("\tThree streams created manually via NewStream")
				t.Print(-1)
            

   Too much trouble for nothing? You can always use Stream.make( 10, 20, 30 ) to
   do this. But notice that this way we can construct our own infinite streams
   easily. Let's make a stream which is an endless series of ones:

				js:

                 function ones() {
                     return new Stream(
                         // the first element of the stream of ones is 1...
                         1,
                         // and the rest of the elements of this stream are given by calling the function ones() (this same function!)
                         ones
                     );
                 }

                 var s = ones();      // now s contains 1, 1, 1, 1, ...
                 s.take( 3 ).print(); // prints 1, 1, 1


				go:


				func Ones(v interface{}, fn STREAMFN) *Stream {
					return NewStream(1, Ones)
				}

				s = NewStream(1, Ones)
				s.Take(3).Print(-1)
            

   Notice that if you use s.print() on an infinite stream, it will print for
   ever, eventually running out of memory. Hence it's best to s.take( n ) before
   you s.print(). Using s.length() on infinite streams is meaningless, so don't
   do it; it will cause an infinite loop (trying to find the end of an endless
   stream). But of course you can use s.map( f ) and s.filter( f ) on infinite
   streams. However, s.walk( f ) will also not run properly on infinite streams.
   So those are some things to keep in mind; make sure you use s.take( n ) if you
   want to take a finite part of an infinite stream.

   Let's see if we can make something more interesting. Here's an alternative and
   interesting way to create the stream of natural numbers:

				js:

                 function ones() {
                     return new Stream( 1, ones );
                 }
                 function naturalNumbers() {
                     return new Stream(
                         // the natural numbers are the stream whose first element is 1...
                         1,
                         function () {
                             // and the rest are the natural numbers all incremented by one
                             // which is obtained by adding the stream of natural numbers...
                             // 1, 2, 3, 4, 5, ...
                             // to the infinite stream of ones...
                             // 1, 1, 1, 1, 1, ...
                             // yielding...
                             // 2, 3, 4, 5, 6, ...
                             // which indeed are the REST of the natural numbers after one
                             return ones().add( naturalNumbers() );
                         }
                     );
                 }
                 naturalNumbers().take( 5 ).print(); // prints 1, 2, 3, 4, 5


				go:

				func NaturalNumbers(v interface{}, fn STREAMFN) *Stream {
					return NewStream(1, func(v interface{}, fn STREAMFN) *Stream {
						return Ones(1, nil).Add(NaturalNumbers(v, fn))
					})
				}

				NaturalNumbers(1, nil).Take(5).Print(-1)

            

   The careful reader will now observe the reason why the second parameter to new
   Stream is a function that returns the tail and not the tail itself. This way
   we can avoid infinite loops by postponing when the tail is evaluated.

   Let's now turn to a little harder example. It's left as an exercise for the
   reader to figure out what the following piece of code does.

				js:

                 function sieve( s ) {
                     var h = s.head();
                     return new Stream( h, function () {
                         return sieve( s.tail().filter( function( x ) {
                             return x % h != 0;
                         } ) );
                     } );
                 }
                 sieve( Stream.range( 2 ) ).take( 10 ).print();


				go:

				/*
				* haskell:
				*  sieve (p : xs) = p : sieve [x | x <- xs, x `mod` p > 0]
				*  take 10 $ sieve [2..]
				*/

				func Sieve(s *Stream) *Stream {
					h := s.Head1()
					return NewStream(h, func (v interface{}, fn STREAMFN) *Stream {
						return Sieve(s.Tail1().Filter(func(x interface{}) bool {
							switch d := x.(type) {
								case int:
									switch dh := h.(type) {
											case int:
												return d % dh != 0
											default:
												return false
									}
								default:
									return false
							}
						}))
					})
				}

				Sieve(Range(2, -1)).Take(10).Print(-1)
            

   Make sure you take some time to figure out what this does. Most programmers
   find it hard to understand unless they have a functional programming
   background, so don't feel bad if you don't get it immediately. Here's a hint:
   Try to find what the head of the printed stream will be. And then try to find
   what the second element of the stream will be (the head of the tail); then the
   third element, and so forth. The name of the function may also help you. If
   you enjoyed this puzzle, here's another.

   If you really can't figure out what it does, just run it and see for yourself!
   It'll be easier to figure out how it does it then.

Ports

   The following ports of stream.js are available:
     * coffeestream is a CoffeeScript port by Michael Blume
     * streamphp is a PHP port by Ryan Gantt
     * python-stream is a Python port by Aris Mikropoulos

Tribute

   Streams aren't in fact a new idea at all. Many functional languages support
   them. The name 'stream' is used in Scheme, a LISP dialect that supports these
   features. Haskell also supports infinite lists. The names 'take', 'tail',
   'head', 'map' and 'filter' are all used in Haskell. A different but similar
   concept also exists in Python and in many other languages; these are called
   "generators".

   These ideas have been around for a long time in the functional programming
   community. However, they're quite new concepts for most Javascript
   programmers, especially those without a functional programming background.

   Many of the examples and ideas come from the book Structure and Interpretation
   of Computer Programs. If you like the ideas here, I highly recommend reading
   it; it's available online for free. It was my inspiration for building this
   library.

   If you prefer a different syntax for streams, you can try out linq.js or, if
   you use node.js, node-lazy may be for you.

Thanks for reading!

   I hope you learned something and that you enjoy using stream.js. I didn't get
   paid to make this library, so if you liked it or it helped you in any way, you
   can buy me a cup of hot chocolate (I don't drink coffee) or just send me a
   e-mail. If you do, make sure to write where you're from and what you do. I
   also enjoy receiving pictures of places from around the world, so feel free to
   attach a picture of yourself in your city!

   Follow @dionyziz

   Fancy donating 3EUR? [ PayPal - The safer, easier way to pay online. ] 
   *

Your rights with stream.js

   The stream.js tutorial is licensed under Creative Commons Attribution 3.0.

   stream.js is licensed under the MIT license.

   Copyright (c)2011 Dionysis Zindros <dionyziz@gmail.com>

   Permission is hereby granted, free of charge, to any person obtaining a copy
   of this software and associated documentation files (the "Software"), to deal
   in the Software without restriction, including without limitation the rights
   to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
   copies of the Software, and to permit persons to whom the Software is
   furnished to do so, subject to the following conditions:

   The above copyright notice and this permission notice shall be included in all
   copies or substantial portions of the Software.

   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
   IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
   FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
   AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
   LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
   OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
   SOFTWARE.

   Some icons are from Default Icon. If you use stream.js I'd love it if you
   e-mailed me with feedback.

   MIT license Creative Commons License
   Lovingly made in London city by dionyziz.
