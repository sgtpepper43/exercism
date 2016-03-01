'use strict';
module.exports = class HelloWorld {
  hello(input) {
    return `Hello, ${input || 'World'}!`;
  }
}
