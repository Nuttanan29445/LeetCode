/**
 * @param {number} n
 * @param {number} k
 * @return {number}
 */
var minChanges = function (n, k) {
  if (n === k) {
    return 0;
  }
  let binN = Number(n).toString(2);
  let binK = Number(k).toString(2);
  if (binK.length < binN.length) {
    binK = "0".repeat(binN.length - binK.length) + binK;
  } else if (binN.length < binK.length) {
    binN = "0".repeat(binK.length - binN.length) + binN;
  }

  let modifiedBinN = "";
  let change = 0;

  for (let i = 0; i < binN.length; i++) {
    if (binN[i] === "1" && binK[i] === "0") {
      modifiedBinN += "0";
      change += 1;
    } else {
      modifiedBinN += binN[i];
    }
  }

  if (modifiedBinN === binK) {
    return change;
  } else {
    return -1;
  }
};

console.log(minChanges(54, 4));
