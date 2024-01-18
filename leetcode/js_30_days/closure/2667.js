/**
 * @return {Function}
 */
var createHelloWorld = function() {
  var msg = "Hello World"
  return function(...args) {
    return msg
  }
};

/**
* const f = createHelloWorld();
* f(); // "Hello World"
*/