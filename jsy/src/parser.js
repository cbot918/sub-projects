function Parser() {
  function A() {
    return "A";
  }

  function B() {
    return "B";
  }

  return {
    A,
    B,
  };
}

export { Parser };
