/**
 * @param {string[]} bank
 * @return {number}
 */
var numberOfBeams = function (bank) {
  let result = 0;
  let consecutiveNonZeroCount = 0;

  for (let sec = 0; sec < bank.length; sec++) {
    const countOnes = (bank[sec].match(/1/g) || []).length;

    if (countOnes !== 0) {
      if (consecutiveNonZeroCount > 0) {
        result += countOnes * consecutiveNonZeroCount;
      }
      consecutiveNonZeroCount = countOnes;
    }
  }

  return result;
};
