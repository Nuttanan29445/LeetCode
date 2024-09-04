/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
var getEncryptedString = function (s, k) {
  let ans = "";
  for (let i = 0; i < s.length; i++) {
    if (i + k < s.length) {
      ans += s[i + k];
    } else {
      ans += s[(i + k) % s.length];
    }
  }
  return ans;
};
