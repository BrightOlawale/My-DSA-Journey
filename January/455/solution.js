/**
 * @param {number[]} g
 * @param {number[]} s
 * @return {number}
 */
var findContentChildren = function (g, s) {
  let contented = 0;
  g.sort((a, b) => a - b);
  s.sort((a, b) => a - b);

  for (const child of g) {
    for (let size = 0; size < s.length; size++) {
      if (child <= s[size]) {
        contented++;
        s.splice(size, 1);
        break;
      }
    }
  }

  return contented;
};
