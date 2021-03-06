package main

const SETTINGS_FILE string = "settings.json"
const WELCOME_BANNER string = "banner_welcome.txt"

const LOG_FOLDER string = "logs"

const GET_PODS string = "kubectl get pods"
const GET_SERVICES string = "kubectl get services"
const GET_NODES string = "kubectl get nodes"
const GET_DEPLOYMENTS string = "kubectl get deployments"
const GET_CONFIG_MAPS string = "kubectl get configmap"
const CURRENT_CONTEXT string = "CURRENT CONTEXT ON KUBERNETS: "
const GET_PERSISTENT_VOLUMES string = "kubectl get pv"
const GET_PERSISTENT_VOLUMES_ARGS string = "--sort-by=.spec.capacity.storage --output=custom-columns=NAME:.metadata.name,CAPACITY:.spec.capacity.storage,STATUS:.status.phase,NAMESPACE:.spec.claimRef.namespace,CLAIM:.spec.claimRef.name,AGE:.metadata.creationTimestamp"

const SHORTCUTS_PODS string = " \n (ENTER) access| (CTRL+Y) return to main menu| (CTRL+D) del| (CTRL+I) describe| (CTRL+L) load| (CTRL+R) refresh"
const SHORTCUTS_SERVICES string = " \n (CTRL+Y) return to main menu| (CTRL+L) load"
const SHORTCUTS_FILTER string = " \n (CTRL+Y) return to main menu"
const SHORTCUTS_COPY_LOGS string = " \n (CTRL+Y) return to main menu| (CTRL+U) return to maintenance menu"
const SHORTCUTS_SETTINGS string = " \n (CTRL+Y) return to main menu"
const SHORTCUTS_MAINTENANCE string = " \n (CTRL+Y) return to main menu"
const SHORTCUTS_NODES string = " \n (CTRL+Y) return to main menu| (CTRL+L) load"
const SHORTCUTS_DEPLOYMENTS string = " \n (CTRL+Y) return to main menu| (CTRL+L) load"
const SHORTCUTS_CONFIG_MAPS string = " \n (CTRL+Y) return to main menu| (CTRL+L) load| (CTRL+I) describe"
const SHORTCUTS_PERSISTENT_VOLUMES string = " \n (CTRL+Y) return to main menu| (CTRL+L) load"
const SHORTCUTS_CONTEXT string = " \n (ENTER) apply context| (CTRL+Y) return to main menu"
const SHORTCUTS_LOAD_CONFIG string = " \n (CTRL+Y) return to main menu| (CTRL+X) return table"
const SHORTCUTS_DESCRIBE string = " \n (CTRL+Y) return to main menu| (CTRL+M) return table"
