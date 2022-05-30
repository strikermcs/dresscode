<template>
    <div class="home_page">
            <v-app-bar color="transparent" flat >
                <ProductTabs :tabs='tabs' @tabChange='ChangeTab'/>
            </v-app-bar>
            <v-main>
            <v-card tile color="#323B50" class="my-2 card_global">
                <v-row>
                    <v-col cols="12" sm="8" style="overflow: scroll; height: 85vh;">
                        <v-window v-model="tab">
                            <v-window-item value="Parasut">
                                <ProductList :products='products'/>
                            </v-window-item>
                            <v-window-item value="Parka">
                                <ProductList :products='parkas'/>
                            </v-window-item>
                        </v-window>
                    </v-col>
                    <v-col cols="12" sm="4">
                        <ThePayment />
                    </v-col>
                </v-row>
            </v-card>
            </v-main> 
    </div>
</template>

<script>
import ThePayment from '../components/ThePayment.vue'
import pic from '../assets/5eb9efce-679f-11eb-b1e0-d8cb8a9c13c3-1-667x1000.jpeg'
import pic2 from '../assets/95718be2-ceb1-11eb-b1e0-d8cb8a9c13c3-21-667x1000.jpeg'
import ProductList from '../components/products/ProductList.vue'
import ProductTabs from '../components/products/ProductTabs.vue'

export default {
    name: 'Home',
    data: () => ({
        tab: null,
        hover: false,

        tabs: [
            {value: 'Parasut', name: 'Parasut',},
            {value: 'Parka', name: 'Parka',},
            {value: 'Steve', name: 'Steve',},
            {value: 'Dennis', name: 'Dennis',}
        ],
        products: [
            {
                avatar: pic,
                prix: '$80',
                title: 'Style four Parasut'
            },

            {
                avatar: pic,
                prix: '$80',
                title: 'Style four Parasut'
            },

            {
                avatar: pic,
                prix: '$80',
                title: 'Style four Parasut'
            },

            {
                avatar: pic,
                prix: '$80',
                title: 'Style four Parasut'
            },

            {
                avatar: pic,
                prix: '$80',
                title: 'Style four Parasut'
            },

            {
                avatar: pic,
                prix: '$80',
                title: 'Style four Parasut'
            }


        ],

        parkas: [
            {
                avatar: pic2,
                prix: '$100',
                title: 'Style four Shoes'
            },

            {
                avatar: pic2,
                prix: '$100',
                title: 'Style four Shoes'
            },

            {
                avatar: pic2,
                prix: '$100',
                title: 'Style four Shoes'
            },

            {
                avatar: pic2,
                prix: '$100',
                title: 'Style four Shoes'
            },

            {
                avatar: pic2,
                prix: '$100',
                title: 'Style four Shoes'
            },

            {
                avatar: pic2,
                prix: '$100',
                title: 'Style four Shoes'
            },
        ]
    }),
    methods: {
        ChangeTab(tab){
            this.tab = tab
        }
    },
    components: {
        ThePayment,
        ProductTabs,
        ProductList
    }
}
</script>


<style>
.v-tab__slider {
    position: absolute;
    bottom: -8px !important;
    left: 50% !important;
    transform: translate(-50%, -50%);
    height: 12px !important;
    width: 20px !important;
    border-radius: 150px 150px 0px 0px;
    border-color: #F34F64;
}


.container {
    padding: 0px !important; 
}

.v-card.card_global {
    min-height: 800px !important;
    width: 100% !important;
}


.v-main {
    padding: 116px 0px 0px 80px !important;
   
}

</style>