export async function getData(url = "") {
  // Default options are marked with *
  const response = await fetch(url, {
    mode: "no-cors", // 'cors' by default
  }).then((data) => {
    console.log(data);
  });
  console.log(response);
  return response.json(); // parses JSON response into native JavaScript objects
}
