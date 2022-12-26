<script>
  import axios from "axios";
  import { Link, navigate } from "svelte-navigator";

  axios
    .post("/logout")
    .then(function (response) {
      console.log(JSON.stringify(response.data));
      alert("Log Out Complete");
      navigate("/");
    })
    .catch(function (error) {
      console.log(error);
    });
</script>
