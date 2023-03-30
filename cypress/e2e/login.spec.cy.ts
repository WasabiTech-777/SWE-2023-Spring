describe('Login Test', () => {
    it('Visits the login page', () => {
      cy.visit('/', { timeout: 5000 })
      cy.url().should('include', 'localhost:4200')
      cy.intercept('GET', '/users',
        (req) => {
          req.continue((res) => {
          expect(res.statusCode).to.be.equal(200)
          })
        }
      )
  
      cy.get('[placeholder="Enter a username"]').type('hello10')
      cy.get('[placeholder="Enter a password"]').type('world10')
  
      cy.get('button').contains('visibility').click()
      cy.get('button').contains('visibility').click()
  
      cy.get('mat-card').find('button').contains('Log In').click()
      cy.intercept('POST', '/login',
        (req) => {
          req.continue((res) => {
          expect(res.statusCode).to.be.equal(200)
          })
        }
      )
      cy.url().should('include', 'profile')
  
    })
  })