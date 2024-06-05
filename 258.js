/**
 * @param {number} num
 * @return {number}
 */
var addDigits = function (num) {
  while (num > 10) {
    let n = 0;
    num = String(num).split("");
    for (let i of num) {
      n += parseInt(i);
    }
    num = n;
  }
  return num;
};

console.log(addDigits(38));
