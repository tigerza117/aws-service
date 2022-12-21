<script>
    import axios from "axios";
    import FormData from 'form-data';
    import {link} from 'svelte-spa-router'
    let acc_id = ''
    let des_acc = ''
	let amount_tra = '' 
    let bank_accounts = [];
    let check_status = 0
    var data = new FormData();
    function handleClick(){
        if (check_status == 0){
            alert('Please Check first')
        }
        else{
            data.append('id', acc_id);
			data.append('acc', des_acc);
			data.append('amount', amount_tra);
            axios.post("https://i-here-ji.tigerza117.xyz/transfer", data, {withCredentials: true}).then(function (response) {
                console.log(JSON.stringify(response.data));
                alert("Transfer Complete");
                acc_id = ''
				des_acc = ''
                amount_tra = ''
                check_status = 0
                reload();
            }).catch(function (error) {
                console.log(error);
            });
        }
        
    }

    function checkState(){
        data.append('id', acc_id);
		data.append('acc', des_acc);
		data.append('amount', amount_tra);
        axios.post("https://i-here-ji.tigerza117.xyz/pre-transfer", data, {withCredentials: true}).then(function (response) {
            console.log(JSON.stringify(response.data));
            alert("Account Confirmed. Please sumbit");
            check_status = 1
        }).catch(function (error) {
            console.log(error);
            data = new FormData()
        });
        
        
    }

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
</script>
<style>
    .ba-card {
    width: 500px;
    border-style: double;
    padding: 10px;
    }
    button:hover{
	background-color: #22d07e;
	color: white;
	cursor: pointer;
}
</style>
<body>
	<h1>Transfer</h1>
    <a href="/Profile" use:link>Back to profile <br><br></a>
	<input bind:value={acc_id} type="text" name="acc_id" placeholder="Your Account ID"/>
	<input bind:value={des_acc} type="text" name="des_acc" placeholder="Destination Account No"/>
	<input bind:value={amount_tra} type="text" name="amount_tra" placeholder="Amount"/>
	<button on:click={checkState}>Check Account</button>
    <button on:click={handleClick}>Submit</button>
    <div class="user-bank-account">
        {#each bank_accounts as bank_account}
        <div class="ba-card">
            <!-- <p>Account</p> -->
            <h4>{bank_account.name}</h4>
            <h4>No.{bank_account.no}</h4>
            <h4>ID: {bank_account.id}</h4>
            <h4>Balance: {bank_account.balance}</h4>
        </div>
        {/each}
    </div>
</body>
