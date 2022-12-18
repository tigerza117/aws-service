<script>
    import {link} from 'svelte-spa-router'
</script>

<style>
    body{
        text-align: center;
        justify-content: center;
    }
</style>
<body>
    <h1> Hello to 3x2 Banking</h1>
    <p>Have an accout? <a href="/login" use:link>Login</a></p>
    <p>Become our member! <a href="/Register" use:link>Register</a></p>
</body>
