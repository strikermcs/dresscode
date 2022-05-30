<template>
    <v-card flat class="card_item pa-2">
        <v-row>
            <v-col cols="12" sm="3" 
            v-for="(product, i) in products" :key="i">
                <Product :product='product'/>
            </v-col>
        </v-row>
    </v-card>
</template>


<script>
import Product from './ProductCard.vue'

export default {
    props:{
        products:{
            required: true
        }
    },
    components:{
        Product
    }
}
</script>

<style scoped>
.card_item {
    background-color: #323B50 !important;
    color: #BABEC9 !important;
    cursor: pointer;
}
</style>