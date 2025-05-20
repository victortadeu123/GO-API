/*async function carregarProdutos() {
  try {
    const resposta = await fetch("http://localhost:8000/products");
    const produtos = await resposta.json();

    const tabela = document.getElementById("lista-produtos");
    tabela.innerHTML = "";

    produtos.forEach(produto => {
      const linha = document.createElement("tr");

      linha.innerHTML = `
        <td>${produto.id}</td>
        <td>${produto.name}</td>
        <td>R$ ${parseFloat(produto.price).toFixed(2)}</td>
        <td>${produto.estoque}</td>
        <td>${produto.categoria}</td>
        <td>${produto.descricao}</td>
        <td>${produto.ativo ? "Sim" : "Não"}</td>
      `;

      tabela.appendChild(linha);
    });
  } catch (erro) {
    console.error("Erro ao carregar produtos:", erro);
  }
}

window.onload = carregarProdutos;*/




document.addEventListener('DOMContentLoaded', function() {
    const productsTable = document.getElementById('lista-produtos');
    const btnCreate = document.getElementById('btnCreateProduct');
    const btnUpdate = document.getElementById('btnUpdateProduct');
    const btnDelete = document.getElementById('btnDeleteProduct');
    


    let selectedProductId = null;


    console.log('ID selecionado:', selectedProductId);
    const id = Number(selectedProductId); 



    loadProducts();

    btnCreate.addEventListener('click', createProduct);
    btnUpdate.addEventListener('click', updateProduct);
    btnDelete.addEventListener('click', deleteProduct);

    async function loadProducts() {
        try {
            const response = await fetch('http://localhost:8000/products');
            if (!response.ok) throw new Error('Erro ao carregar produtos');
            const products = await response.json();
            renderProducts(products);
        } catch (error) {
            console.error(error);
            alert(error.message);
        }
    }

    function renderProducts(products) {
    productsTable.innerHTML = '';
   
    products.forEach(product => {
        const row = document.createElement('tr');
        row.dataset.id = product.id;

        row.innerHTML = `
            <td>${product.id}</td>
            <td>${product.name}</td>
            <td>R$ ${product.price.toFixed(2)}</td>
            <td>${product.estoque}</td>
            <td>${product.categoria}</td>
            <td>${product.descricao || '-'}</td>
            <td>${product.ativo ? 'Sim' : 'Não'}</td>
        `;

        // Evento de clique na linha
        row.addEventListener('click', function () {
            // Remove classe 'selected' de todas as linhas
            document.querySelectorAll('#lista-produtos tr').forEach(r => {
                r.classList.remove('selected');
            });

            // Adiciona a classe 'selected' à linha atual
            this.classList.add('selected');

            // Salva o ID do produto selecionado
            selectedProductId = this.dataset.id;

            console.log('Produto selecionado ID:', selectedProductId); // <-- debug
        });

        productsTable.appendChild(row);
    });
}


    async function createProduct() {
        try {
            const name = prompt('Nome do produto:');
            if (!name) return;

            const price = parseFloat(prompt('Preço do produto:'));
            if (isNaN(price)) return;

            const estoque = parseInt(prompt('Estoque:'), 10);
            if (isNaN(estoque)) return;

            const categoria = prompt('Categoria:');
            const descricao = prompt('Descrição:');
            const ativo = confirm('Produto está ativo?');

            const newProduct = {
                name,
                price,
                estoque,
                categoria,
                descricao,
                ativo
            };

            const response = await fetch('http://localhost:8000/products', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(newProduct)
            });

            if (!response.ok) throw new Error('Erro ao criar produto');

            alert('Produto criado com sucesso');
            loadProducts();
        } catch (error) {
            console.error(error);
            alert(error.message);
        }
    }

    async function updateProduct() {
        if (!selectedProductId) {
            alert('Selecione um produto para atualizar');
            return;
        }

        try {
            const product = await getProduct(selectedProductId);

            const name = prompt('Nome do produto:', product.name);
            const price = parseFloat(prompt('Preço do produto:', product.price));
            const estoque = parseInt(prompt('Estoque:', product.estoque), 10);
            const categoria = prompt('Categoria:', product.categoria);
            const descricao = prompt('Descrição:', product.descricao || '');
            const ativo = confirm('Produto está ativo?');

            const updatedProduct = {
                name,
                price,
                estoque,
                categoria,
                descricao,
                ativo
            };

            const response = await fetch(`http://localhost:8000/products/${selectedProductId}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(updatedProduct)
            });

            if (!response.ok) throw new Error('Erro ao atualizar produto');

            alert('Produto atualizado com sucesso');
            loadProducts();
        } catch (error) {
            console.error(error);
            alert(error.message);
        }
    }
    

    async function deleteProduct() {
        if (!selectedProductId) {
            alert('Selecione um produto para deletar');
            return;
        }

        if (!confirm('Deseja realmente deletar o produto?')) return;

        try {
            const response = await fetch(`http://localhost:8000/products/${selectedProductId}`, {
                method: 'DELETE'
            });

            if (!response.ok) throw new Error('Erro ao deletar produto');

            alert('Produto deletado com sucesso');
            selectedProductId = null;
            loadProducts();
        } catch (error) {
            console.error(error);
            alert(error.message);
        }
    }

    async function getProduct(id) {
        const response = await fetch(`http://localhost:8000/products/${id}`);
        if (!response.ok) throw new Error('Produto não encontrado');
        return await response.json();
    }

    

    
});



