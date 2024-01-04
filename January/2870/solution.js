/**
 * @param {number[]} nums
 * @return {number}
 */
var minOperations = function (nums) {
  let hashMap = new Map();
  let output = 0;
  let prevNum = 0;
  nums.sort((a, b) => a - b);
  nums.push(0);

  for (const num of nums) {
    if (!hashMap.has(num)) {
      hashMap.set(num, 1);
    } else {
      hashMap.set(num, hashMap.get(num) + 1);
    }

    if ((prevNum && prevNum != num) || num == 0) {
      let occ = hashMap.get(prevNum);

      if (occ === 1) {
        return -1;
      }
      if (occ % 3 !== 0) {
        if (occ === 2 || occ === 4) {
          output += occ / 2;
        } else {
          let new_occ = occ - 2;
          if (new_occ % 3 !== 0) {
            occ = occ - 4;
            output += 2;
          } else {
            occ = new_occ;
            output += 1;
          }
          output += occ / 3;
        }
      } else {
        output += occ / 3;
      }
    }
    prevNum = num;
  }

  return output;
};
