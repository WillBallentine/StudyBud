using System;
using Microsoft.EntityFrameworkCore;
using StudyBud.Data.Interfaces;
using StudyBud.Models;


namespace StudyBud.Data
{
	public class GeneralUserDAL : IGeneralUserDAL
	{
		ApplicationDbContext _dbContext;

		public GeneralUserDAL(ApplicationDbContext dbContext)
		{
			_dbContext = dbContext;
		}

		public async void NewUserAsync(string Id, string phone, string fName, string lName)
		{
			if (await _dbContext.Users.FindAsync(Id) is User found)
			{
				found.PhoneNumber = phone;
				found.FName = fName;
				found.LName = lName;

				await _dbContext.SaveChangesAsync();
			}
		}
	}
}

