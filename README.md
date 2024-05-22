# Ethereum blockchain parser

[TOC]

## 1. Goal

build an ethereum blockchain parser that will allow to query transactions for subscribed addresses

Users not able to receive push notifications for incoming/outgoing transactions. By Implementing Parser interface we
would be able to hook this up to notifications service to notify about any incoming/outgoing transactions.

## 2. Usage

* ```get_current_block``` return the latest block number
* ```subscribe (eth address)``` subscribe to an address
* ```transaction (eth address)``` show transactions from an subscribed address
* ```exit``` exit system

## 3. Examples

when you start the parser you will see the following output:
>test address:0xae2fc483527b8ef99eb5d9b44875f005ba1fae13

and this address is for testing only

then you can execute cmds from previews section 
like :
>subscribe 0xae2fc483527b8ef99eb5d9b44875f005ba1fae13

and then

>transaction 0xae2fc483527b8ef99eb5d9b44875f005ba1fae13 
>
you will be able to see transaction result



