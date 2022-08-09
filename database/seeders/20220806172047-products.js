'use strict';

module.exports = {
  async up (queryInterface, Sequelize) {
    const products = [];
    for (let i = 1; i < 100; i++) {
      products.push({
        name: `Name #${i}`,
        description: `Describing item #${i}`,
        price: i,
        createdAt: new Date().getTime() / 1000,
        updatedAt: new Date().getTime() / 1000
      })
    }

    await queryInterface.bulkInsert('products', products);
  },

  async down (queryInterface, Sequelize) {
    /**
     * Add commands to revert seed here.
     *
     * Example:
     * await queryInterface.bulkDelete('People', null, {});
     */
  }
};
