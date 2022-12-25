<script>
    import axios from "axios";
    import FormData from 'form-data';
    // import {link} from 'svelte-spa-router'
    import NavBarComp from "../components/NavBarComp.svelte";
    import Footer from "../components/FooterComp.svelte";

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

<div style="overflow: hidden; background-color: black; max-width:100%; width:100%; max-height:100vh;">
        
    <!-- navbar -->
    <nav class="w-full h-[10%] border-gray-200 px-2 sm:px-4 py-2.5 rounded bg-gray-900">
        <div class="container flex flex-wrap items-center justify-between mx-auto">
            <a href="https://flowbite.com/" class="flex items-center">
                <img
                    src="https://flowbite.com/docs/images/logo.svg"
                    class="h-6 mr-3 sm:h-9"
                    alt="Flowbite Logo"
                />
                <span
                    class="self-center text-xl font-semibold whitespace-nowrap text-white"
                    >3X2 Banking</span
                >
            </a>
            <div class="hidden w-full md:block md:w-auto" id="navbar-default">
                <ul
                    class="flex flex-col p-4 mt-4 border rounded-lg md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0"
                >
                    <NavBarComp 
                        pagelink="/profile" 
                        pagename="Profile" 
                        home={false} 
                    />
                    <NavBarComp
                        pagelink="/transfer"
                        pagename="Transfer"
                        home={true}
                    />
                    <NavBarComp
                        pagelink="/deposit"
                        pagename="Deposit"
                        home={false}
                    />
                    <NavBarComp
                        pagelink="/login"
                        pagename="Logout"
                        home={false}
                    />
                </ul>
            </div>
        </div>
    </nav>

    <div class="flex max h-auto bg-gray-500">
        <div class="w-1/4 h-1/4">
            <div class="overflow-y-auto py-4 px-3 bg-gray-800 h-screen">
                <div style="justify-content: center; align-item:center; display:flex; padding: 3%; ">
                    <h1 class="pt-10 mb-2 text-center text-7xl font-bold tracking-tight text-gray-900 dark:text-white">Transfer</h1>
                </div>
                <ul class="space-y-6 pt-16">
                        <div class="">
                            <label for="id" class="block mb-2 text-xl font-medium text-white">Enter Your Account-ID:</label>
                            <input
                            class="border text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 bg-gray-600 border-gray-500 text-white"
                                bind:value={acc_id}
                                type="text"
                                name="acc_id"
                                placeholder="Your Account-ID"
                            />
                        </div>
                        <div class="">
                            <label for="No" class="block mb-2 text-xl font-medium text-white">Enter The Destination Account Number:</label>
                            <input
                            class="border text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 bg-gray-600 border-gray-500 text-white"
                                bind:value={des_acc}
                                type="text"
                                name="des_acc"
                                placeholder="Destination Account No."
                            />
                        </div>
                        <div class="">
                            <label for="balance" class="block mb-2 text-xl font-medium text-white">Enter Amount Number:</label>
                            <input
                            class="border text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 bg-gray-600 border-gray-500 text-white"
                                bind:value={amount_tra}
                                type="text"
                                name="amount_tra" 
                                placeholder="Amount"
                            />
                        </div>
                        <div class="space-y-2">
                            <button class="text-white w-full focus:ring-4 focus:outline-none font-medium rounded-lg text-sm px-5 py-2.5 text-center bg-blue-600 hover:bg-blue-700 focus:ring-blue-800" on:click={checkState}>Check Account</button>
                        </div>
                        <div class="space-y-2">
                            <button class="text-white w-full focus:ring-4 focus:outline-none font-medium rounded-lg text-sm px-5 py-2.5 text-center bg-blue-600 hover:bg-blue-700 focus:ring-blue-800" on:click={handleClick}>Submit Transfer</button>
                        </div>
                </ul>
            </div>
        </div>
        <div class="w-3/4 bg-gray-800">
            <div class="">
                <div class="flow-root overflow-y-auto h-[41.5rem] w-full">
                    {#each bank_accounts as bank_account}
                    <ul class="divide-y divide-gray-700 bg-gray-700" style="padding-left: 1%; border: white; border-style:groove;">
                        <li class="py-3 sm:py-4">
                            <div class="flex items-center space-x-4">
                                <div class="flex-1 min-w-0 w-3/6">
                                    <div class="flex">
                                        <div class="pr-2">
                                            <p class="text-6xl font-medium truncate text-white">
                                            {bank_account.name}
                                            </p>
                                        </div>
                                        <div class="flex justify-center items-end">
                                            <p class="text-2xl truncate text-gray-400">ID: {bank_account.id}</p>
                                        </div>
                                    </div>
                                    <br>
                                    <p class="text-3xl truncate text-gray-400">
                                        No.{bank_account.no} 
                                    </p>
                                    <br><p class="text-sm truncate text-gray-400">
                                        Account Created: {bank_account.created_at}
                                    </p>
                                    <p class="text-sm truncate text-gray-400">
                                        Account Updated: {bank_account.updated_at}
                                    </p>
                                </div>
                                <div class="pr-5 column-flex justify-end text-base font-semibold text-white w-3/6">
                                    <h1 class="text-4xl text-right">Balance: </h1>
                                    <h1 class="text-right text-8xl">${bank_account.balance}</h1>
                                </div>
                            </div>
                        </li>
                    </ul>
                    {/each}
                </div>

            </div>
        </div>
    </div>

    <Footer/>
</div>