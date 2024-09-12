/**
 * @param {string} allowed
 * @param {string[]} words
 * @return {number}
 */
var countConsistentStrings = function (allowed, words) {
  let ans = 0;
  words.forEach((word) => {
    let res = word;
    for (let i = 0; i < allowed.length; i++) {
      res = res.replaceAll(allowed[i], "");
    }
    if (res == "") {
      ans += 1;
    }
  });
  return ans;
};
