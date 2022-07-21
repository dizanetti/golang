package main

const SETTINGS_FILE string = "settings.json"
const WELCOME_BANNER string = "banner_welcome.txt"

const GET_PODS string = "kubectl get pods"
const GET_SERVICES string = "kubectl get services"
const GET_NODES string = "kubectl get nodes"
const GET_DEPLOYMENTS string = "kubectl get deployments"
const CURRENT_CONTEXT string = "CURRENT CONTEXT ON KUBERNETS: "
const GET_PERSISTENT_VOLUMES string = "kubectl get pv"
const GET_PERSISTENT_VOLUMES_ARGS string = "--sort-by=.spec.capacity.storage --output=custom-columns=NAME:.metadata.name,CAPACITY:.spec.capacity.storage,STATUS:.status.phase,NAMESPACE:.spec.claimRef.namespace,CLAIM:.spec.claimRef.name,AGE:.metadata.creationTimestamp"

const SHORTCUTS_PODS string = " \n (ENTER) access| (CTRL+Y) return menu| (CTRL+V) return table| (CTRL+D) delete| (CTRL+I) describe| (CTRL+R) refresh"
const SHORTCUTS_SERVICES string = " \n (CTRL+Y) return to Menu"
const SHORTCUTS_FILTER string = " \n (CTRL+Y) return to Menu"
const SHORTCUTS_SETTINGS string = " \n (CTRL+Y) return to Menu"
const SHORTCUTS_NODES string = " \n (CTRL+Y) return to Menu"
const SHORTCUTS_DEPLOYMENTS string = " \n (CTRL+Y) return to Menu"
const SHORTCUTS_PERSISTENT_VOLUMES string = " \n (CTRL+Y) return to Menu"
const SHORTCUTS_CONTEXT string = " \n (ENTER) apply context| (CTRL+Y) return to Menu"
