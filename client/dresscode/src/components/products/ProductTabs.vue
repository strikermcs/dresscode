<template>
    <v-tabs color="#F34F64" v-model="tabActive" @click="tabsHandler()">
        <v-tab class="withoutupercase" 
        v-for="(tab, i) in tabs" :key="i"
        :value="tab.value">{{tab.name}}</v-tab>    
    </v-tabs>
</template>

<script>
export default {
    props: {
        tabs: {
            required: true
        }
    },
    data:() => ({
        tabActive: null
    }),
    methods:{
        tabsHandler(){
           this.$emit('tabChange', this.tabActive)
        }
    }
}
</script>

<style scoped>
.v-tab.withoutupercase {
    text-transform: none !important;
    color: #778290;
}
</style>