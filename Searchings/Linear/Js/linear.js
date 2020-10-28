function search(key, arr) {
    for (let i = 0; i< arr.length; i++) {
      if (arr[i] === key){
        console.log("index =", i)
      }
    }
  }
  
  search(3, [1,4,5,3]); 