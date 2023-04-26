// Reverse Polish Notation
//
// 40 2 + -> 42
// 2 3 - -> -1
// 5 4 * 2 / -> 10
// 5 4 2 * - -> -3
//
// + - * /
class RPNCalc {
  constuctor() {

  }

  setInput(input) {
    this.input = input;
  }

  tokenize() {
    this.tokens = this.input.split(/\s+/);
  }

  parseFirst() {
    var tokenSet = {
      o1: Number(this.tokens[0]),
      o2: Number(this.tokens[1]),
      op: this.tokens[2]
    };
    this.tokens = this.tokens.slice(3);
    return tokenSet;
  }

  parseNext() {
    var tokenSet = {
      o2: Number(this.tokens[0]),
      op: this.tokens[1]
    };
    this.tokens = this.tokens.slice(2);
    return tokenSet;
  }

  doOne(o1, o2, op) {
    var result;
    if (op == '+') {
      result = o1 + o2;
    } else if (op == '-') {
      result = o1 - o2;
    } else if (op == '*') {
      result = o1 * o2;
    } else if (op == '/') {
      result = o1 / o2;
    }

    return result;
  }

  enter() {
    var {o1, o2, op} = this.parseFirst();
    var result = this.doOne(o1, o2, op);
    while (this.tokens.length) {
      var {o2, op} = this.parseNext();
      result = this.doOne(result, o2, op);
    }

    return result;
  }

  go(input) {
    this.setInput(input);
    rpnCalc.tokenize();
    var result = rpnCalc.enter();
    console.log(result);
  }
}

var rpnCalc = new RPNCalc();
// rpnCalc.go('0 1 + 2 / 2 * 2 + 1 + 1 - 1 +');
// rpnCalc.go('2 2 + 2 +');
// rpnCalc.go('2 2 -');
// rpnCalc.go('2 2 *');
// rpnCalc.go('2 2 /');
// rpnCalc.go('40 2 +');
// rpnCalc.go('2 3 -');
// rpnCalc.go('5 4 * 2 /');
// rpnCalc.go('5 4 2 * -');


class RPNCalc2 {
  constructor() {
  }

  go(input) {
    var result = '';
    while(true) {
      var tokenSet = input.match(/\d+\s+\d+\s+\D/);
      if (!tokenSet) {
        return result;
      }
      var match = tokenSet[0];
      var index = tokenSet.index;

      var [o1, o2, op] = match.split(/\s/);
      console.log(o1, o2, op);

      o1 = Number(o1);
      o2 = Number(o2);

      result = eval(`${o1} ${op} ${o2}`);

      if (input.length == match.legth) {
        return result;
      }

      input = `${input.substring(0, index)} ${result} ${input.substring(index + match.length)}`;
      console.log(input);
    }

    return result;
  }
}


var rpnCalc2 = new RPNCalc2();
// rpnCalc2.go('40 2 +');
// rpnCalc2.go('2 3 -');
// rpnCalc2.go('5 4 * 2 /');
rpnCalc2.go('5 4 2 * -');


