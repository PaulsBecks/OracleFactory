// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

interface PubSubOracle {
    function subscribeInteger(string calldata topic, address smartContract) external;
    function unsubscribeInteger(string calldata topic, address smartContract) external;
}

interface IntegerCallback {
    function integerCallback(string calldata topic, uint256 value) external;
}

contract Oracle is PubSubOracle {

    struct Subscription {
        address smartContract;
        bool active;
    }

    struct SubscriptionManager {
        Subscription[] subscriptions;
        bool initialized;
        mapping (address => uint) positions;
    }

    mapping (string => uint256) public storedIntegers;
    mapping (string => SubscriptionManager) private integerSubscriptionManagers;
    mapping (string => string) public storedStrings;

    event Subscribe(string topic, address smartContract);
    event Unsubscribe(string topic, address smartContract);
    event PublishInteger(string topic, address smartContract, uint256 value);

    function publishInteger(string memory topic, uint256 value) public {
        storedIntegers[topic] = value;
        // notify subscriptions
        SubscriptionManager storage subscriptionManager = integerSubscriptionManagers[topic];
        for (uint i = 0; i < subscriptionManager.subscriptions.length; i++){
            Subscription storage subscription = subscriptionManager.subscriptions[i];
            if(subscription.active){
                IntegerCallback(subscription.smartContract).integerCallback(topic, value);
            }
        }
    }

    function subscribeInteger(string memory topic, address smartContract) public {
        SubscriptionManager storage subscriptionManager = integerSubscriptionManagers[topic];
        // make sure to start at 1 to use 0 as empty value in map
        subscriptionManager.initialized = true;
        uint new_length = subscriptionManager.subscriptions.length + 1;
        uint position = subscriptionManager.positions[smartContract];
        if (position != 0) {
            // already subscribed
            subscriptionManager.subscriptions[position-1].active = true;
            emit Subscribe(topic,  smartContract);
            return;
        }
        subscriptionManager.subscriptions.push(Subscription(smartContract, true));
        subscriptionManager.positions[smartContract] = new_length;
        emit Subscribe(topic,  smartContract);
    }

    function unsubscribeInteger(string memory topic, address smartContract) public {
        SubscriptionManager storage subscriptionManager = integerSubscriptionManagers[topic];
        uint index = subscriptionManager.positions[smartContract];
        // check if index exists
        if (index == 0) {
            // string(abi.encodePacked("No integer subscription for topic ", topic)));
            return;
        }
        subscriptionManager.subscriptions[index-1].active = false;
        emit Unsubscribe(topic, smartContract);
    }
}

