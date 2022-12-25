<script>
    import axios from "axios";
    import {push} from 'svelte-spa-router'

    axios.post('https://i-here-ji.tigerza117.xyz/logout').then(function (response) {
        console.log(JSON.stringify(response.data));
        alert("Log Out Complete")
        push('/');
    })
    .catch(function (error) {
    console.log(error);
    });
</script>