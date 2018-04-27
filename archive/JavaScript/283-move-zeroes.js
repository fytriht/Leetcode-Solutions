let moveZeroes = nums => {
  let j = 0;
  nums.forEach((n, i) => {
    if (n != 0) {
      [nums[i], nums[j]] = [nums[j], nums[i]];
      j++;
    }
  });
};
