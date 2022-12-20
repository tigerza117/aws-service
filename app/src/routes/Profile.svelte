<script>
    import axios from "axios";
    import NotFound from "./NotFound.svelte";
    let name = '';
    let email = '';
    let userid = '';
    let idcreated_at = '';
    let idupdated_at = '';
    let bank_accounts = [
        {
        "balance": 100,
        "created_at": "2022-12-14T21:19:05.98+07:00",
        "id": 5,
        "name": "yeah_account_1",
        "no": "6908675654",
        "updated_at": "2022-12-14T21:19:05.98+07:00"
        },  
        {
        "balance": 2000,
        "created_at": "2022-12-14T21:19:24.919+07:00",
        "id": 7,
        "name": "yeah_acc_2",
        "no": "8510122835",
        "updated_at": "2022-12-14T21:19:24.919+07:00"
        }
    ];

    var config = {
        method: 'get',
        url: 'https://i-here-ji.tigerza117.xyz/profile',
        withCredential: true,
        headers: { }
    };

    axios.get('https://i-here-ji.tigerza117.xyz/profile').then(function (response) {
        if (response.data) {
            console.log(JSON.stringify(response.data));
            console.log(response.data.name)
            name = response.data.name;
            email = response.data.email;
            userid = response.data.id;
            idcreated_at = response.data.created_at;
            idupdated_at = response.data.updated_at;
            // bank_accounts = response.data.account;
        } else {
            console.log('No data received from the server');
    }
    }).catch(function (error) {
        console.log(error);
    });
</script>

<body>
    <section>
        <h1>INFORMATION</h1>
        <h2>Name: {name}</h2>
        <h2>Email: {email}</h2>
        <h2>UserID: {userid}</h2>
        <h4>ID Created: {idcreated_at}</h4>
        <h4>ID Updated: {idupdated_at}</h4>
    </section>
    <div class="user-bank-account">
        {#each bank_accounts as bank_account}
        <div class="ba-card">
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
    }

</style>