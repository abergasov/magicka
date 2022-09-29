<template>
  <div class="hello">
    <h1 v-if="account">Hello, {{ account }}</h1>
    <h1 v-else>anonymous</h1>
    <button v-if="!account" @click="connectWallet">Connect wallet</button>
    <h3 v-if="ethBalance">ETH: {{ethBalance}}</h3>
    <h3 v-if="daiBalance">DAI: {{daiBalance}}</h3>
    <h3 v-if="usdcBalance">USDC: {{usdcBalance}}</h3>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type {AxiosInstance} from 'axios'
import Web3 from 'web3';
import { BlockHeader, Block } from 'web3-eth'
import { AbiItem } from 'web3-utils'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance
  }
}

export default defineComponent({
  name: 'HelloWorld',
  props: {
    msg: String,
  },
  data() {
    return {
      account: '',
      ethBalance: '',
      daiBalance: '',
      usdcBalance: '',
    }
  },
  computed: {
    metamaskExist() :boolean {
      return typeof window.ethereum !== 'undefined';
    },
  },
  methods: {
    async connectWallet() {
      if (!this.metamaskExist) {
        alert('Please install metamask');
        return;
      }
      const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
      this.account = accounts[0];
      const contractsResponse = await this.$axios.get('/contracts');
      const web3 = new Web3(process.env.VUE_APP_GETH_IPC || 'wss://localhost:8545');

      const ethBalance = await web3.eth.getBalance(this.account);
      this.ethBalance = web3.utils.fromWei(ethBalance, "ether");
      
      const batch = new web3.BatchRequest();
      const contractDAI = new web3.eth.Contract(this.getABI() as AbiItem[], contractsResponse.data.DAI)
      const contractUSDC = new web3.eth.Contract(this.getABI() as AbiItem[], contractsResponse.data.USDC)
      // batch.add(contractDAI.methods.balanceOf(this.account).call().then((res: string) => {
      //   this.daiBalance = web3.utils.fromWei(res, "ether");
      // }));

      batch.add(contractDAI.methods.balanceOf(this.account).call().then((res: string) => {
        this.daiBalance = web3.utils.fromWei(res, "ether");
      }));
      batch.add(contractUSDC.methods.balanceOf(this.account).call().then((res: string) => {
        this.daiBalance = web3.utils.fromWei(res, "ether");
      }));


      // for (let token of contractsResponse.data) { // tokens is list of erc20 tokens
      //
      //   batch.add(contractDAI.methods.balance(address).call.request({from: '0x0000000000000000000000000000000000000000'}, callback2));
      //
      // }

      //batch.execute()

      await contractDAI.methods.balanceOf(this.account).call().then((res: string) => {
        this.daiBalance = web3.utils.fromWei(res, "ether");
      });
      await contractUSDC.methods.balanceOf(this.account).call().then((res: string) => {
        this.usdcBalance = web3.utils.fromWei(res, "ether");
      });
    },
    getABI() {
      return [
        {
          "constant": true,
          "inputs": [
            {
              "name": "_owner",
              "type": "address"
            }
          ],
          "name": "balanceOf",
          "outputs": [
            {
              "name": "balance",
              "type": "uint256"
            }
          ],
          "payable": false,
          "stateMutability": "view",
          "type": "function"
        },
      ];
    }
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
