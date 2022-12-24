<script>
	import axios from "axios";
	import {link, push} from 'svelte-spa-router'
	let email = "";
	let password = "";
	function handleClick() {
	const data = { email: email, password: password };
	  axios.post("https://i-here-ji.tigerza117.xyz/login", data, {withCredentials: true}).then(result => {
			alert("Login Complete");
			push('/profile');
	  }).catch(err => {
		alert("Invalid Username or password");
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
	
	<h1 class="text-3xl font-bold underline">Log in</h1>
	<input bind:value={email} type="email" name="email" placeholder="Email"/>
	<input bind:value={password} type="password" name="password" placeholder="Password"/>
	<button on:click={handleClick}>Submit</button>
	<p>Don't have an account? <a href="/Register" use:link>Register</a></p>
</body>
