# Introduction

Imagine every time you receive an email from your grandma you want to play [Glorious from Macklemore](https://www.youtube.com/watch?v=7OrLroFa0AI) via spotify. You could either check your emails every minute, and when it arrives go to your spotify account and play the song. This would be a little stressfull, so you could automate this process with a workflow automation tool.


## Workflow Automation

To automate professional or private tasks some people like to write computer programs. Not everyone is able to code, so companies created no-code workflow automation platforms, like [Zapier](https://www.zapier.com), [n8n](https://www.n8n.io) or [IFTTT](https://www.ifttt.com). These platforms let you combine services like Google Spreadsheets, Spotify, mail programs or messengers services, to automate your tasks in so called workflows.

These workflows propagate data from triggers, that provide the data, through brokers, to data consumers as shown in the picture below. 

![](./Provider-Broker-Consumer.png)

In our use case Google Mail could be the provider, that alerts the workflow that an email arrived, and spotify would be the consumer.

## The Problem

Many services can be integrated as described above. But, some services are deployed to a special environment called blockchain. These services are called smart contracts. You can think of them like regular computer programs. One challenge Smart contracts face, is that they cannot communicate on their own to other computer programs, that are not deployed to their blockchain. We call services that live inside a blockchain on-chain and those that are not deployed inside a blockchain off-chain.

There is one on-chain service called Crypto Kitties where you can breed, buy and sell digital cats. Let's assume you want to play Glorious by Macklemore every time a new cat is born. In this case the crypto kitties event could be a provider in an automation workflow. Well, this not yet possible in the way you could when receiving an email. 

This means, integrating smart contracts into workflow automation is difficult. But other computer programs, called oracles, can create a bridge between smart contract and off-chain programs. There are two kind of oracles that are particularly interesting to us. Push-based inbound oracles and push-based outbound oracles, which I will call for simplicity just inbound and outbound from now on. 

Inbound Subscriptions receive some data from an off-chain program and will forward this data to a on-chain smart contract service, as displayed in the graphic below.

Outbound Subscription go the other way they listen to events created by smart-contracts and will forward the data to an off-chain program, also shown in the graphic below.

![](Inbound-Outbound-Subscription.png)

You can see, that the oracle always contains a provider and a consumer. If we want to use a smart contract as a provider we need to create an outbound oracle. If the smart contract should be a consumer, we create an inbound oracle.

In our example this could look like this: a new cat is born in the Crypto Kitties smart contract -> this is registered by a smart contract provider of an outbound oracle -> web service consumer sends it to workflow -> workflow contacts spotify to play the song.

But also the other way around could be possible. Where a smart contract acts as a consumer. For example let's say you want to transfer a cat to your brother when he sends you an email. Then you could use GMail as a provider and an inbound oracle as the consumer.

## The Software
The software you will test, will try to let no-coders create these oracles, so no-coders can integrate smart contracts into their workflows.

## The Tasks

We ask you to do some tasks. Before starting please open [our tool](https://oracles.work)


### Using a smart contract as provider
The first part will be about using the kitties birth event as a provider in a workflow.

Go to Zapier https://zapier.com and log in with the credentials provided and navigate to the zaps.
1. The Zap "Test 1" is a workflow to start a spotify song. Create an outbound oracle, that provides this workflow with the crypto kitties birth event. 

### Using a smart contract as consumer
1. The Zap "Test 2", create an inbound oracle that consumes the events of this worflow.


### Creating Providers and Consumers
1. In zap Test 1, what would you need to change to let the outbound oracle you created in step 6 be used as a provider in this workflow

1. Create an account and connect it to the Ethereum
blockchain with the following credentials:
    * Ethereum Private Key: 0x34567ab1289efb761298732.
    * Ethereum Gateway: wss://infura.com/ws/v3/hfjdksafehuka
1. Create a smart contract provider for the
crypto kitties birth event.
    * Smart Contract Address: 0x06012c8cf97BEaD5deAe237070F9587f8E7A266d
    * Contract Address Synonym: Crypto Kitties
    * Event Name: Birth
    * Parameters:
        owner of type address,  kittyId of type uint256, matronId of type uint256 ,  sireId of type uint256 and genes of type uint256
1. Create a smart contract consumer for
the given smart method:
    * Smart Contract Address: 0x06012c8cf97BEaD5deAe237070F9587f8E7A266d
    * Contract Address Synonym: Crypto Kitties
    * Method Name: Transfer
    * Parameters:
        _to of type address, _tokenId of type uint256
1. Create a web service provider to trigger transfering an kitty
1. Create a web service consumer with
that will publish data to a workflow to play Glorious on Spotify.
    * URL: https://hooks.zapier.com/hooks/catch/1796279/b1239873

