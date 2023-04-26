// Next Smallest Numeric Palindrome
//
// bao@ujet.co
//
// to create a numeric plaindrome, we modify digits in the left and right halves of the number so that they 'mirror' each other.
// to make sure the output is larger than the input, we add one to the initial number, otherwise we would return the number itself (if it were already a palindrome).
//
// moving from least significant digit (lsd) to most significant digit (msd)...
//   the lsd of the mirrored pair is always incremented.
//   if the digit goes 'past zero', a carry is added to the next msd guaranteeing that the number increases monotonically.
//   if the number has an even number of digits, the carry propagates from the lsd of the mirrored pair to the msd.
//   after this 'cross-over' the digits are already mirrored, so we continue processing to...
//      1. process to the 'carry' and
//      2. mirror changes caused by the carry back to the lsds.
//   if the number has an odd number of digits...
//      the middle digit is handled as a degnerate case.
//      it only changes in the event of a carry. if it is a 9, it propagates the carry.
//
function nsp(start) {
  // we must be at least one larger...
  start++;

  var arr = String(start).split('').map((char) => Number(char)).reverse();
  var len = arr.length;
  var iLast = len - 1;

  var carry = 0;
  for (var i = 0; i < len; i++) {
    // add the carry
    if (carry) {
      if (arr[i] == 9) {
        arr[i] = 0;
      } else {
        arr[i]++;
        carry = 0;
      }
    }

    // working on the least significant digits?
    if (i < len / 2) {
      // we could also carry if the msd < lsd
      carry = (arr[iLast - i] < arr[i]) ? 1 : carry;

      // mirror the most-significant digit to the least significant digit
      arr[i] = arr[iLast - i];
    }
    else {
      // mirror the most-significant digit to the least significant digit
      arr[iLast - i] = arr[i];
    }
  }

  return Number(arr.join(''));
}


var {expect} = require('chai');
expect(nsp(0)).to.equal(1);
expect(nsp(1)).to.equal(2);
expect(nsp(2)).to.equal(3);
expect(nsp(3)).to.equal(4);
expect(nsp(4)).to.equal(5);
expect(nsp(5)).to.equal(6);
expect(nsp(6)).to.equal(7);
expect(nsp(7)).to.equal(8);
expect(nsp(8)).to.equal(9);
expect(nsp(9)).to.equal(11);
expect(nsp(10)).to.equal(11);
expect(nsp(11)).to.equal(22);
expect(nsp(12)).to.equal(22);
expect(nsp(13)).to.equal(22);
expect(nsp(14)).to.equal(22);
expect(nsp(15)).to.equal(22);
expect(nsp(16)).to.equal(22);
expect(nsp(17)).to.equal(22);
expect(nsp(18)).to.equal(22);
expect(nsp(19)).to.equal(22);
expect(nsp(20)).to.equal(22);
expect(nsp(21)).to.equal(22);
expect(nsp(22)).to.equal(33);
expect(nsp(90)).to.equal(99);
expect(nsp(99)).to.equal(101);
expect(nsp(101)).to.equal(111);
expect(nsp(120)).to.equal(121);
expect(nsp(121)).to.equal(131);
expect(nsp(190)).to.equal(191);
expect(nsp(191)).to.equal(202);
expect(nsp(919)).to.equal(929);
expect(nsp(999)).to.equal(1001);
expect(nsp(99998)).to.equal(99999);
expect(nsp(99999)).to.equal(100001);
expect(nsp(123456)).to.equal(124421);
expect(nsp(123321)).to.equal(124421);
expect(nsp(199999)).to.equal(200002);
expect(nsp(12345678)).to.equal(12355321);

console.log('finished');