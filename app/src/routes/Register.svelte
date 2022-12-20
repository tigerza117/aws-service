<script>
    import axios from "axios";
    import {link} from 'svelte-spa-router'
    let name = "";
    let email = "";
    let password = "";
    
    function handleClick() {
    const data = { name: name, email: email, password: password };
      axios.post("https://i-here-ji.tigerza117.xyz/register", data).then(result => {
            alert("Register Complete");
            name = '';
			email = '';
			password = '';
		}).catch(err => {
        console.log(err);
      });
    }
</script>
<style>
body{
    text-align: center;
    justify-content: center;
}
button:hover{
	background-color: #22d07e;
	color: white;
	cursor: pointer;
}
</style>
<body>
    <h1>Register</h1>
    <input bind:value={name} type="name" name="name" placeholder="Name"/>
    <input bind:value={email} type="email" name="email" placeholder="Email"/>
    <input bind:value={password} type="password" name="password" placeholder="Password"/>
    <button on:click={handleClick}>Submit</button>
    <p>Have an account? <a href="/login" use:link>Login</a></p>
</body>