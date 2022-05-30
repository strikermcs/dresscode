<template>
     <v-hover v-slot="{isHovering, props}" close-delay="200">
        <v-card :elevation="isHovering ? 16 : 2" 
        class="mx-auto" height="420" tile color="#333C51" aling="center"
        v-bind="props"
        >
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn variant="contained" size="small" :color="isHovering ? '#F34F64' : '#777D8B'">
                    {{product.prix}}
                </v-btn>
            </v-card-actions>
            <v-img :src="product.avatar" height="320" cover/>
            <v-card-text :class="isHovering ? 'red--text' : 'white--text'" class="text-center">
                {{product.title}}
            </v-card-text>
        </v-card>
    </v-hover>
</template>


<script>
export default {
    props: {
        product :{
            required: true
        }
    }
}
</script>

<style scoped>
.red--text {
    color: #F34F64;
}

.white--text {
    color: #BABEC9;
}
</style>