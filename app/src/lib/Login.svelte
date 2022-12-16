<script>
    import axios from "axios";
    let name = "";
    let email = "world";
    let password = "";
    const data = { name: name, email: email, password: password };
    function handleClick() {
        axios.post("https://i-here-ji.tigerza117.xyz/register", data).then(result => {
            if (result.message === "OK") {
                alert("Yes");
            } else {
                alert("No");
            }
        });
    }
</script>

<input bind:value={name} type="name" name="name"/>
<input bind:value={email} type="email" name="email"/>
<input bind:value={password} type="password" name="password"/>
<button on:click={handleClick}>Submit</button>
<h1>Log in</h1>