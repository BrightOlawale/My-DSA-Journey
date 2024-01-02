/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var findMatrix = function (nums) {
  let hashMap = new Map();
  let output = [];

  for (const num of nums) {
    if (!hashMap.has(num)) {
      hashMap.set(num, 0);
    } else {
      hashMap.set(num, hashMap.get(num) + 1);
    }

    const index = hashMap.get(num);
    if (!output[index]) {
      output[index] = [num];
    } else {
      output[index].push(num);
    }
  }

  return output;
};
