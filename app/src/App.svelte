<script>
	import Router from 'svelte-spa-router'
	import Home from './routes/Home.svelte'
	import Login from './routes/Login.svelte'
	import NotFound from './routes/NotFound.svelte'
	import Register from './routes/Register.svelte'
	import Profile from './routes/Profile.svelte'
	import Transfer from './routes/Transfer.svelte';
	import Deposit from './routes/Deposit.svelte';
	import axios from 'axios'
	const routes = {
		'/':Home,
		'/login':Login,
		'/register': Register,
		'/profile': Profile,
		'/transfer': Transfer,
		'/deposit': Deposit,
		'*': NotFound,
	}
	axios.defaults.withCredentials = true
</script>

<Router {routes}/>


