describe('Full End to End Test', () => {
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
      cy.reload()
      cy.get('mat-card').should('include.text','Articles Typed: 0')
      cy.get('mat-toolbar').should('include.text','Welcome')
      cy.get('mat-toolbar').find('button').contains('Articles').click()
      cy.get('mat-card').contains('Japan').click()
      cy.get('[placeholder="Start Typing"]').type('The first people in Japan were the Ainu people and other Jōmon people. They were closer related to Europeans or Mongols. They were later conquered and replaced by the Yayoi people (early Japanese and Ryukyuans). The Yayoi were an ancient ethnic group that migrated to the Japanese archipelago mainly from southeastern China during the Yayoi period (300 BCE–300 CE). Modern Japanese people have primarily Yayoi ancestry at an average of 97%. The indigenous Ryukyuan and Ainu peoples have more Jōmon ancestry on the other hand.')
      cy.get('[placeholder="infoBar"]').should("have.text", "Text Length: 525 || # of Correct Keystrokes: 525 || # of Mistakes: 0")
      cy.get('[placeholder="WPM"]').should("have.text", "Words Per Minute: 175")
    })
  })