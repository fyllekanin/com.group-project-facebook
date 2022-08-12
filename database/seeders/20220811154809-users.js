'use strict';

module.exports = {
  async up (queryInterface, Sequelize) {
    await queryInterface.bulkInsert('users', [
        {
		username: 'test',
		password: '$2a$10$R5NfXusiOvVFG.uDStz65.KjHo2eeuSol3qRdNNhLbwE9K1dJp8tS',
		createdAt: new Date().getTime(),
		updatedAt: new Date().getTime()
	}
    ]);

    /**
     * Add seed commands here.
     *
     * Example:
     * await queryInterface.bulkInsert('People', [{
     *   name: 'John Doe',
     *   isBetaMember: false
     * }], {});
    */
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
