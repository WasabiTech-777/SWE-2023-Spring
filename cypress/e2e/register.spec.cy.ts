describe('My First Test', () => {
  it('Visits the initial project page', () => {
    cy.visit('/', { timeout: 5000 })
    cy.url().should('include', 'localhost:4200')
    cy.intercept('GET', '/users',
      (req) => {
        req.continue((res) => {
        expect(res.statusCode).to.be.equal(200)
        })
      }
    )

    cy.get('button').contains('Register').click()
    cy.url().should('include', 'register')

    cy.get('[placeholder="Enter a username"]').type('Alligator4')
    cy.get('[placeholder="Enter a password"]').type('@l!G@T0R2023')

    cy.get('button').contains('visibility').click()
    cy.get('button').contains('visibility').click()

    cy.get('button').contains('Sign').click()
    cy.intercept('POST', '/users',
      (req) => {
        req.continue((res) => {
        expect(res.statusCode).to.be.equal(200)
        })
      }
    )

  })
})
