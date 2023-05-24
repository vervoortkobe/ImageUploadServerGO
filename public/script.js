const parent = document.getElementById("posts");
fetch("http://localhost:3000/all")
.then(res => {
  return res.json();
})
.then(data => {
  data.json.forEach(e => {
    const newP = document.createElement("p");
    newP.innerHTML = `${e.name} - ${e.id}`;
    const newImg = document.createElement("img");
    newImg.src = `http://localhost:3000/${e.id}`;
    parent.appendChild(newP);
    parent.appendChild(newImg);
  });
})
.catch(err => {
  console.log(err);
});