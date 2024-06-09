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

		public void NewUser(string userId, string firstName, string lastName, string email, string phone)
		{
			var user = new User { UserId = userId, FName = firstName, LName = lastName, Name = firstName + " " + lastName, Email = email, Phone = phone };
			_dbContext.Add(user);
			_dbContext.SaveChanges();
		}
	}
}

