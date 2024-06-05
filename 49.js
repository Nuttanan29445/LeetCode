/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function (strs) {
  let z = [];
  let d = {};
  let ans = [];
  for (let i of strs) {
    const temp = i.split("").sort().join("");
    if (!d[temp]) {
      d[temp] = [];
      z.push(temp);
    }
    d[temp].push(i);
  }
  for (let j of z) {
    ans.push(d[j]);
  }
  return ans.sort((a, b) => a.length - b.length);
};

const input = ["eat", "tea", "tan", "ate", "nat", "bat"];
console.log(groupAnagrams(input));
