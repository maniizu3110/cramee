export default function({ $axios }) {
  $axios.onRequest(config => {
    config.headers.common["Authorization"] =
      "Bearer " + localStorage.getItem("Bearer");
  });
}
