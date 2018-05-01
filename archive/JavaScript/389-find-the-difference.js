let findTheDifference = (s, t) =>
  String.fromCharCode(
    (s + t)
      .split('')
      .map(b => b.charCodeAt() - 'a'.charCodeAt())
      .reduce((ret, b) => (ret ^= b), 0) + 'a'.charCodeAt(),
  )
