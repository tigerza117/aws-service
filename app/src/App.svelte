<script>
  import Foot from "./components/Foot.svelte";
  import { Route, Router } from "svelte-navigator";
  import Login from "./routes/Login.svelte";
  import Register from "./routes/Register.svelte";
  import Profile from "./routes/Profile.svelte";
  import Transfer from "./routes/Transfer.svelte";
  import Deposit from "./routes/Deposit.svelte";
  import Navbar from "./components/Navbar.svelte";
  import { isLogin } from "./stores/auth";
  import PrivateRoute from "./routes/PrivateRoute.svelte";
  import PublicRoute from "./routes/PublicRoute.svelte";


</script>
<link
  href="https://fonts.googleapis.com/icon?family=Material+Icons"
  rel="stylesheet"
/>
<div class="h-full min-h-screen min-w-full">
  <Navbar />
  <div class="mx-auto min-h-screen w-full bg-gray-200">
    <Router>
      <PublicRoute path="/auth/*">
        <Route path="login" component="{Login}" />
        <Route path="register" component="{Register}" />
      </PublicRoute>

      <PrivateRoute path="/*">
        <Route path="" component="{Profile}" />
        <Route path="transfer" component="{Transfer}" />
        <Route path="deposit" component="{Deposit}" />
      </PrivateRoute>
    </Router>
  </div>

  <Foot />
</div>
