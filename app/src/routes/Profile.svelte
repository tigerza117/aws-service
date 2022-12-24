<script>
    import NavBarComp from "../components/NavBarComp.svelte";
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
<!-- navbar -->
<nav class="bg-white border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900">
    <div class="container flex flex-wrap items-center justify-between mx-auto">
      <a href="https://flowbite.com/" class="flex items-center">
          <img src="https://flowbite.com/docs/images/logo.svg" class="h-6 mr-3 sm:h-9" alt="Flowbite Logo" />
          <span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white">Flowbite</span>
      </a>
      <button data-collapse-toggle="navbar-default" type="button" class="inline-flex items-center p-2 ml-3 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600" aria-controls="navbar-default" aria-expanded="false">
        <span class="sr-only">Open main menu</span>
        <svg class="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd"></path></svg>
      </button>
      <div class="hidden w-full md:block md:w-auto" id="navbar-default">
        <ul class="flex flex-col p-4 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
          <NavBarComp pagelink="/" pagename="Home" home={true}/>
          <NavBarComp pagelink="/transfer" pagename="Transfer" home={false}/>
          <NavBarComp pagelink="/deposit" pagename="Deposit" home={false}/>
          <NavBarComp pagelink="/login" pagename="Logout" home={false}/>
        </ul>
      </div>
    </div>
  </nav>
<!-- navbar dark:bg-gray-900 -->

    <!-- <a href="/transfer" use:link>Transfer</a> -->
    <!-- <a href="/deposit" use:link>Deposit</a> -->
    <!-- <a href="/login" use:link>Logout</a> -->

    <div class="flex">
        <div class="w-1/4 h-fit">
            <section>
                <p class="text-5xl">INFORMATION</p>
                <h2 class="text-2xl">Name: {name}</h2>
                <h2>Email: {email}</h2>
                <h2>UserID: {userid}</h2>
                <h4>ID Created: {idcreated_at}</h4>
                <h4>ID Updated: {idupdated_at}</h4>
            </section>
        </div>
        <div class="w-3/4">
            <div class="">
                <section>
                    <h1>Create New Bank Account</h1>
                    <input bind:value={new_account} type="new_account" name="new_account" placeholder=" Account Name"/>
                    <button on:click={handleClick}>Create</button>
                </section>
            </div>
            <div class="">
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
            </div>
        </div>

    </div>







<!-- footer -->
<footer class="fixed bottom-0 left-0 z-20 p-4 w-full bg-white border-t border-gray-200 shadow md:flex md:items-center md:justify-between md:p-6 dark:bg-gray-800 dark:border-gray-600">
    <span class="text-sm text-gray-500 sm:text-center dark:text-gray-400">© 2022 <a href="https://flowbite.com/" class="hover:underline">Flowbite™</a>. All Rights Reserved.
    </span>
    <ul class="flex flex-wrap items-center mt-3 text-sm text-gray-500 dark:text-gray-400 sm:mt-0">
    </ul>
</footer>
<!-- footer -->
</body>

<style>
	section {
        text-align: center;
        justify-content: center;
		width: 100%;
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
