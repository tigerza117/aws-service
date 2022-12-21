<script>
    import axios from "axios";
    import {link} from 'svelte-spa-router'
    let name = '';
    let new_account = '';
    let email = '';
    let userid = '';
    let idcreated_at = '';
    let idupdated_at = '';
    let bank_accounts = [];

    axios.get('https://i-here-ji.tigerza117.xyz/profile').then(function (response) {
        if (response.data) {
            console.log(JSON.stringify(response.data));
            console.log(response.data.name)
            name = response.data.name;
            email = response.data.email;
            userid = response.data.id;
            idcreated_at = response.data.created_at;
            idupdated_at = response.data.updated_at;
            bank_accounts = response.data.account;
        } else {
            console.log('No data received from the server');
    }
    }).catch(function (error) {
        console.log(error);
    });

    function reload(){
        axios.get('https://i-here-ji.tigerza117.xyz/accounts').then(function (response) {
            if (response.data) {
                console.log(JSON.stringify(response.data));
                console.log(response.data.name)
                bank_accounts = response.data;
                
            } else {
                console.log('No data received from the server');
        }
        }).catch(function (error) {
            console.log(error);
        });
    }
    reload();
    

    // /*Create Bank Account*/
    function handleClick() {
    const data = {name: new_account};
      axios.put("https://i-here-ji.tigerza117.xyz/account", data).then(result => {
            alert("Create Complete");
            new_account = '';
            reload();
		}).catch(err => {
        console.log(err)
        alert(err);
      });
    }


</script>

<body>
    <a href="/transfer" use:link>Transfer</a>
    <a href="/deposit" use:link>Deposit</a>
    <a href="/login" use:link>Logout</a>
    <section>
        <h1>INFORMATION</h1>
        <h2>Name: {name}</h2>
        <h2>Email: {email}</h2>
        <h2>UserID: {userid}</h2>
        <h4>ID Created: {idcreated_at}</h4>
        <h4>ID Updated: {idupdated_at}</h4>
    </section>

    <section>
        <h1>Create New Bank Account</h1>
        <input bind:value={new_account} type="new_account" name="new_account" placeholder="Account Name"/>
        <button on:click={handleClick}>Create</button>
    </section>

    <div class="user-bank-account">
        {#each bank_accounts as bank_account}
        <div class="ba-card">
            <!-- <p>Account</p> -->
            <h1>{bank_account.name}</h1>
            <h3>No.{bank_account.no}</h3>
            <h3>ID: {bank_account.id}</h3>
            <h2>Balance: {bank_account.balance}</h2>
            <h4>Account Created: {bank_account.created_at}</h4>
            <h4>Account Updated: {bank_account.updated_at}</h4>
        </div>
        {/each}
    </div>
</body>

<style>
	section {
        text-align: center;
        justify-content: center;
		width: 350px;
		display: grid;
		grid-template-columns: 1fr 1fr;
		padding: 10px;
		box-shadow: 2px 2px 4px #dedede;
		border: 1px solid #888;
		margin-bottom: 10px;
        display: flex;
        flex-direction: column;
	}
    .ba-card {
        width: 500px;
        border-style: double;
        padding: 10px;
    }

</style>
