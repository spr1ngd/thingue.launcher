import {defineStore} from 'pinia'
import InstancePanel from "@/panels/InstancePanel.vue";
import ClientPanel from '@/panels/ClientPanel.vue'
import CloudResPanel from '@/panels/CloudResPanel.vue'

const panelNameComponentMap = {
    "clientPanel": ClientPanel,
    "instancePanel": InstancePanel,
    "cloudResPanel": CloudResPanel,
}

export const usePanelStore = defineStore('panel', {
    state: () => {
        return {
            open: false,
            name: "",
            data: {},
            width: 300
        }
    },
    getters: {
        component: (state) => {
            if (state.name) {
                // return import(`@/panels/${state.name}.vue`)
                return panelNameComponentMap[state.name]
            } else {
                return null;
            }
        },
    },
    actions: {
        togglePanel(name, data, width = 300) {
            if (name !== this.name || data !== this.data) {
                if (this.open) {
                    this.name = name
                    this.data = data
                    this.width = width
                } else {
                    this.openPanel(name, data, width)
                }
            } else {
                if (this.open) {
                    this.closePanel()
                } else {
                    this.openPanel(name, data, width)
                }
            }
        },
        openPanel(name, data, width = 300) {
            this.open = true
            this.name = name
            this.data = data
            this.width = width
        },
        closePanel() {
            this.open = false
            this.name = ""
            this.data = {}
        }
    }
})