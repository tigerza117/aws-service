import './app.css'
import App from './App.svelte'
import axios from "axios";

axios.defaults.baseURL = "https://dmce1m4ypodn6.cloudfront.net/api"

const app = new App({
  target: document.getElementById('app')
})

export default app
