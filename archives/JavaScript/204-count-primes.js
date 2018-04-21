let countPrimes = n => {
  let m = Array(n).fill(true)
  for (let i = 2; i * i < n; i++) {
    if (m[i]) for (let j = i * i; j < n; j += i) m[j] = false
  }
  return m.slice(2).reduce((cnt, isPrime) => (isPrime ? cnt + 1 : cnt), 0)
}
